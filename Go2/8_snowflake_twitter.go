/*

	分布式id生成器

有时我们需要能够⽣成类似 MySQL 自增 ID 这样不断增大，同时又不会重复的 ID。以支持业务中的⾼并发场景。
比较典型的是电商促销时短时间内会有大量的订单涌入到系统，比如每秒 10w+。明星出轨时会有大量热情的粉丝发微博以表心意，同样会在短时间内产生大量的消息。

分布式ID生成，就我来看主要是2个流派，各有利弊，没有完美的实现。

1，snowflake流派。

它用于twitter的微博ID，因为是timeline按发布时间排序，所以这个算法是用毫秒时间戳作为ID的左半部，从而可以实现按时间有序。
像新浪微博也是在使用类似的ID生成算法，snowflake的好处是去中心化，但是依赖时钟的准确性，最差的情况是时钟发生了回退，那么ID就会重复；
而如果不开启NTP同步时钟，那么不同节点分配的时间不同，也会影响feed流的排序，所以在我看来只能说基本可用，一旦时钟回退比较大的区间，服务是完全不可用的。
美团在这方面做了一些工作，主要还是在发现回退以及报警方面的事情，可以参考： Leaf — 美团点评分布式ID生成系统 。

2，mysql流派。

该流派使用广泛，基本原理就是mysql的自增主键。最初为了扩展性能，会通过部署多台mysql，为每个mysql设置不同的起始id，从而实现横向扩展性。
mysql支持自定义表的auto_increment属性，可以用于控制起始ID：


在插入数据库之前，我们需要给这些消息、订单先打上一个 ID，然后再插⼊到我们的数据库。
对这个 ID 的要求是希望其中能带有一些时间信息，这样即使我们后端的系统对消息进行了分库分表，也能够以时间顺序对这些消息进⾏排序。

Twitter 的 snowflake 算法(https://github.com/bwmarrin/snowflake)是这种场景下的一个典型解法。
先来看看 snowflake 是怎么回事，如下图所示：

+--------------------------------------------------------------------------+
| 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
+--------------------------------------------------------------------------+


http://c.biancheng.net/uploads/allimg/190828/4-1ZRQ3110KJ.gif

⾸先确定我们的数值是 64 位的 int64 类型，被划分为了四部分，不含开头的第一个 bit，因为这个 bit 是符号位。
用 41 位来表示收到请求时的时间戳，单位为毫秒，
然后五位来表示数据中心的 ID，
然后再五位来表示机器的实例 ID，
最后是 12 位的循环自增 ID（到达 1111,1111,1111 后会归 0）。

这样的机制可以⽀持我们在同一台机器上，同一毫秒内产⽣ 2 ^ 12 = 4096 条消息。一秒共 409.6 万条消息。从值域上来讲完全够⽤了。
数据中心加上实例 ID 共有 10 位，可以⽀持我们每数据中心部署 32 台机器，所有数据中心共 1024 台实例。
表示 timestamp 的 41 位，可以⽀持我们使用 69 年。当然，我们的时间毫秒计数不会真的从 1970 年开始记，
那样我们的系统跑到 2039/9/7 23:47:35 就不能用了，所以这里的 timestamp 实际上只是相对于某个时间的增量，
比如我们的系统上线是 2018-08-01，那么我们可以把这个 timestamp 当作是从 2018-08-01 00:00:00.000 的偏移量。


worker_id 分配
timestamp、datacenter_id、worker_id 和 sequence_id 这四个字段中，timestamp 和 sequence_id 是由程序在运⾏期⽣成的。
但 datacenter_id 和 worker_id 需要我们在部署阶段就能够获取得到，并且一旦程序启动之后，就是不可更改的了（想想，如果可以随意更改，可能被不慎修改，造成最终生成的 ID 有冲突）。
一般不同数据中⼼的机器，会提供对应的获取数据中心 ID 的 API，所以 datacenter_id 我们可以在部署阶段轻松地获取到。
而 worker_id 是我们逻辑上给机器分配的一个 ID，这个要怎么办呢？比较简单的想法是由能够提供这种自增 ID 功能的工具来支持，比如 MySQL:

从 MySQL 中获取到 worker_id 之后，就把这个 worker_id 直接持久化到本地，以避免每次上线时都需要获取新的 worker_id，让单实例的 worker_id 可以始终保持不变。
当然，使用 MySQL 相当于给我们简单的 id 生成服务增加了一个外部依赖，依赖越多，我们的服务的可运维性就越差。

考虑到集群中即使有单个 ID 生成服务的实例挂了，也就是损失一段时间的一部分 ID，
所以我们也可以更简单暴力一些，把 worker_id 直接写在 worker 的配置中，上线时由部署脚本完成 worker_id 字段替换。

*/

package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/snowflake"
)

func main() {

	// Create a new Node with a Node number of 1

	node, err := snowflake.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	for i := 0; i < 3; i++ {
		// Generate a snowflake ID.
		id := node.Generate()

		// Print out the ID in a few different ways.
		fmt.Printf("Int64  ID: %d\n", id)
		fmt.Printf("String ID: %s\n", id)
		fmt.Printf("Base2  ID: %s\n", id.Base2())
		fmt.Printf("Base64 ID: %s\n", id.Base64())

		// Print out the ID's timestamp
		fmt.Printf("ID Time  : %d\n", id.Time())

		// Print out the ID's node number
		fmt.Printf("ID Node  : %d\n", id.Node())

		// Print out the ID's sequence number
		fmt.Printf("ID Step  : %d\n", id.Step())

		// Generate and print, all in one.
		fmt.Printf("ID       : %d\n", node.Generate().Int64())

		fmt.Println(
			"\n",
		)
	}
}

/*

输出：

id 1428247897438162944
node:  1 step:  0 time:  1629355808969

id 1428247897442357248
node:  1 step:  0 time:  1629355808970

id 1428247897442357249
node:  1 step:  1 time:  1629355808970

*/
