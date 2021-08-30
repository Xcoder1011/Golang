package main

import (
	"fmt"
	"os"
	"sync"
)

/*
	defer（延迟执行语句）

在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行

也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。

当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）


*/
func main() {

	fmt.Println("defer begin")

	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(3)
	fmt.Println("defer end")

	/// 延迟调用是在 defer 所在函数结束时进行，函数结束可以是正常返回时，也可以是发生宕机时。

	/*

		打印：

			defer begin
			defer end
			3
			2
			1

	*/

	// 处理业务或逻辑中涉及成对的操作，比如打开和关闭文件、接收请求和回复请求、加锁和解锁等
	// 最容易忽略的就是在每个函数退出处正确地释放和关闭资源。
	// defer 语句正好是在函数退出时执行的语句，所以使用 defer 能非常方便地处理资源释放问题。

}

/// 1) 使用延迟并发解锁
func readValue(key string) int {

	// 为防止竞态问题，使用 sync.Mutex 进行加锁
	valueByKeyGuard.Lock() // 使用互斥量加锁

	v := valueByKey[key]

	valueByKeyGuard.Unlock() // 解锁。

	return v
}

/// 使用 defer 语句对上面的语句进行简化

func readValueDefer(key string) int {

	// 为防止竞态问题，使用 sync.Mutex 进行加锁
	valueByKeyGuard.Lock() // 使用互斥量加锁

	v := valueByKey[key]

	// defer后面的语句不会马上调用, 而是延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()

	return v
}

var (
	// 一个演示用的映射
	// map 默认不是并发安全的，准备一个 sync.Mutex 互斥量保护 map 的访问。
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

/// 2) 使用延迟释放文件句柄
/// 文件的操作需要经过打开文件、获取和操作文件资源、关闭资源几个过程，
/// 如果在操作完毕后不关闭文件资源，进程将一直无法释放文件资源，

// 根据文件名查询其大小
func fileSize(filename string) int64 {
	// 根据文件名打开文件, 返回文件句柄和错误
	f, err := os.Open(filename)
	// 如果打开时发生错误, 返回文件大小为0
	if err != nil {
		return 0
	}
	// 取文件状态信息
	info, err := f.Stat()

	// 如果获取信息时发生错误, 关闭文件并返回文件大小为0
	if err != nil {
		f.Close()
		return 0
	}
	// 取文件大小
	size := info.Size()
	// 关闭文件
	f.Close()

	// 返回文件大小
	return size
}

/// 使用 defer 对代码进行简化
func fileSizeDefer(filename string) int64 {
	f, err := os.Open(filename)
	// 不能将 defer f.Close() 代码放在此处处，一旦文件打开错误，f 将为空，在延迟语句触发时，将触发宕机错误。
	if err != nil {
		return 0
	}
	// 延迟调用Close, 此时Close不会被调用
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		// defer机制触发, 调用Close关闭文件
		return 0
	}
	size := info.Size()
	// defer机制触发, 调用Close关闭文件
	return size
}
