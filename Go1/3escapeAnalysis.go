package main

import (
	"fmt"
)

/*
	变量逃逸（Escape Analysis）—— 自动决定变量分配方式，提高运行效率
	命令行执行： go run -gcflags "-m -l" 3escapeAnalysis.go

	# command-line-arguments
	./3escapeAnalysis.go:23:13: ... argument does not escape
	./3escapeAnalysis.go:23:13: a escapes to heap
	./3escapeAnalysis.go:23:22: dummy(12) escapes to heap
	0 12
*/
func main() {

	// var a int

	// void()

	// fmt.Println(a, dummy(12)) // 变量 a 和 dummy(12)  逃逸到堆

	fmt.Println(dummy2()) // 打印 &{}
}

// 1. 逃逸分析

func void() {

}

func dummy(b int) int {

	// 变量 c 是整型，其值通过 dummy() 的返回值“逃出”了 dummy() 函数
	var c int

	c = b
	// 变量 c 的值被复制并作为 dummy() 函数的返回值返回，
	// 即使变量 c 在 dummy() 函数中分配的内存被释放，
	// 也不会影响 main() 中使用 dummy() 返回的值。变量 c 使用栈分配不会影响结果。
	return c
}

// 2. 取地址发生逃逸

/*
	# command-line-arguments
	./3escapeAnalysis.go:53:6: moved to heap: c   // 将 c 移到堆中。
	./3escapeAnalysis.go:25:13: ... argument does not escape
	&{}
*/

type Data struct {
}

func dummy2() *Data {
	// 实例化c为Data类型
	var c Data
	// 返回函数局部变量地址
	// Go 编译器已经确认如果将变量 c 分配在栈上是无法保证程序最终结果的，
	// 如果这样做，dummy2() 函数的返回值将是一个不可预知的内存地址
	return &c

}

// 总结：

// 编译器觉得变量应该分配在堆和栈上的原则是：
// 变量是否被取地址；
// 变量是否发生逃逸。
