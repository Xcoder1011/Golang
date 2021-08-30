package main

import (
	"fmt"
	"runtime"
)

/*
	宕机（panic）——-- 程序终止运行


Go语言的类型系统会在编译时捕获很多错误，但有些错误只能在运行时检查，如数组访问越界、空指针引用等，这些运行时错误会引起宕机。
一般而言，当宕机发生时，程序会中断运行，并立即执行在该 goroutine（可以先理解成线程）中被延迟的函数（defer 机制），
随后，程序崩溃并输出日志信息，日志信息包括 panic value 和函数调用的堆栈跟踪信息，panic value 通常是某种错误信息。

对于每个 goroutine，日志信息中都会有与之相对的，发生 panic 时的函数调用堆栈跟踪信息，
我们不需要再次运行程序去定位问题，日志信息已经提供了足够的诊断依据，因此，在我们填写问题报告时，一般会将宕机和日志信息一并记录。
panic 一般用于严重错误，对于大部分漏洞，我们应该使用Go语言提供的错误机制，而不是 panic。

panic() 的声明为为：  func panic(v interface{})    //panic() 的参数可以是任意类型的。

*/

/*
	宕机恢复（recover）——-- 防止程序崩溃

Recover 是一个Go语言的内建函数，可以让进入宕机流程中的 goroutine 恢复过来，recover 仅在延迟函数 defer 中有效
在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果
在其他语言里，宕机往往以异常的形式存在，底层抛出异常，上层逻辑通过 try/catch 机制捕获异常，
没有被捕获的严重异常会导致宕机，捕获的异常可以被忽略，让代码继续运行。
Go语言没有异常系统，其使用 panic 触发宕机类似于其他语言的抛出异常，recover 的宕机恢复机制就对应其他语言中的 try/catch 机制。
*/

func main() {

	/// 1) 手动触发宕机
	// panicTest()

	/// 2) 让程序在崩溃时继续执行
	RecoverTest()

}

/// 1) 手动触发宕机
func panicTest() {

	// Go语言可以在程序中手动触发宕机，让程序崩溃，这样开发者可以及时地发现错误，同时减少可能的损失。
	// Go语言程序在宕机时，会将堆栈和 goroutine 信息输出到控制台，

	// 当 panic() 触发的宕机发生时，panic() 后面的代码将不会被运行
	// 但是在 panic() 函数前面已经运行过的 defer 语句依然会在宕机发生时发生作用，
	// 这个特性可以用来在宕机发生前进行宕机信息处理。

	defer fmt.Println("宕机后要做的事情1")
	defer fmt.Println("宕机后要做的事情2")

	panic("手动触发宕机")

	fmt.Println("panic() 后面的代码将不会被运行")

	/*

		打印：


		宕机后要做的事情2
		宕机后要做的事情1
		panic: 手动触发宕机

		goroutine 1 [running]:
		main.panicTest()
			/Users/.../Go/src/Go4/2_panic.go:49 +0xdb
		main.main()
			/Users/.../Go/src/Go4/2_panic.go:32 +0x25
		exit status 2

	*/
}

/// 2) 让程序在崩溃时继续执行

func RecoverTest() {

	fmt.Println("运行前")

	// 允许一段手动触发的错误

	ProtectRun(func() {
		fmt.Println("手动宕机前")

		panic(&panicContext{"手动触发panic"})

		fmt.Println("手动宕机后")

	})

	// 故意造成空指针访问错误

	ProtectRun(func() {

		fmt.Println("赋值宕机前")

		var a *int

		*a = 1

		fmt.Println("赋值宕机后")

	})

	fmt.Println("运行后")

	/*

		打印：

			运行前
			手动宕机前
			error: &{手动触发panic}
			赋值宕机前
			runtime error: runtime error: invalid memory address or nil pointer dereference
			运行后

	*/

}

// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string // 所在函数
}

// 保护方式允许一个函数
func ProtectRun(entry func()) {

	// 延迟处理的函数
	defer func() {

		// 发生宕机时，获取panic传递的上下文并打印
		// recover() 获取到 panic 传入的参数。

		err := recover()

		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}

	}()

	entry()
}
