package main

import (
	"fmt"
)

/*
	闭包（Closure）——-- 引用了外部变量的匿名函数


 函数 + 引用环境 = 闭包

 Go语言中闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，
 即使已经离开了自由变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量

一个函数类型就像结构体一样，可以被实例化，函数本身不存储任何信息，只有与引用环境结合后形成的闭包才具有“记忆性”
函数是编译期静态的概念，而闭包是运行期动态的概念。

*/
func main() {

	/// 1. 在闭包内部修改引用的变量

	/// 闭包对它作用域上部的变量可以进行修改，修改引用的变量会对变量进行实际修改，
	str := "hello world"
	// 创建一个匿名函数
	foo := func() {

		// 在匿名函数中并没有定义 str，str 的定义在匿名函数之前
		// 此时，str 就被引用到了匿名函数中形成了闭包。
		str = "hello foo"

		fmt.Println(str) // hello foo
	}

	// 调用匿名函数
	foo()

	/// 2. 闭包的记忆效应

	// 创建一个累加器, 初始值为1
	accumulator := Accumulate(1) // 返回的 accumulator 是类型为 func()int 的函数变量。
	// 被捕获到闭包中的变量让闭包本身拥有了记忆效应
	fmt.Println(accumulator()) // 2
	// 闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命期一直存在，
	fmt.Println(accumulator()) // 3
	// 闭包本身就如同变量一样拥有了记忆效应。
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator) // 0xc0000b0020

	// 创建一个累加器, 初始值为10
	accumulator2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2()) // 11
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2) // 0xc0000b0028

	// accumulator 与 accumulator2 输出的函数地址不同，因此它们是两个不同的闭包实例。
	// 每调用一次 accumulator 都会自动对引用的变量进行累加。

	/// 3. 闭包实现生成器

	// 闭包的记忆效应被用于实现类似于设计模式中工厂模式的生成器，
	// 下面的例子展示了创建一个玩家生成器的过程。

	// 创建一个玩家生成器
	generator := playerGenerate("zhangsan")
	// 返回玩家的名字和血量
	name, hp := generator()

	fmt.Println(name, hp) // 打印 zhangsan 150

}

// 返回一个为初始值创建的闭包函数。
func Accumulate(num int) func() int {
	// 返回一个闭包
	// 每次返回会创建一个新的函数实例。
	return func() int {
		num++
		return num
	}
}

// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGenerate(name string) func() (string, int) {
	// 闭包还具有一定的封装性，hp变量是 playerGenerate 的局部变量，外部无法直接访问及修改这个变量
	hp := 150 // 血量一直为150
	// 返回创建的闭包
	return func() (string, int) {
		// 将变量引用到闭包中
		return name, hp
	}
}

/*

其它编程语言中的闭包

闭包（Closure）在某些编程语言中也被称为 Lambda 表达式。
闭包对环境中变量的引用过程也可以被称为“捕获”，在 C++11 标准中，捕获有两种类型，分别是引用和复制，
可以改变引用的原值叫做“引用捕获”，捕获的过程值被复制到闭包中使用叫做“复制捕获”。

在 Lua 语言中，将被捕获的变量起了一个名字叫做 Upvalue，因为捕获过程总是对闭包上方定义过的自由变量进行引用。

C++ 与 C# 中为闭包创建了一个类，而被捕获的变量在编译时放到类中的成员中，
闭包在访问被捕获的变量时，实际上访问的是闭包隐藏类的成员。
*/
