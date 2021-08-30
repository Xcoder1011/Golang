package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hello World!")

	sinceTime()

	swap(10, 20)

}

// 计算函数执行时间
func sinceTime() {
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}

	elapsed := time.Since(start)
	// elapsed := time.Now().Sub(start)    // time.Now().Sub() 的功能类似于 Since() 函数

	fmt.Println("该函数执行完成耗时：", elapsed) // 该函数执行完成耗时： 29.90367ms
}

/*

基于堆栈式的程序执行模型决定了函数是语言的一个核心元素
1) 函数调用规约
Go语言函数使用的是 caller-save 的模式，即由调用者负责保存寄存器，所以在函数的头尾不会出现push ebp; mov esp ebp这样的代码，
相反其是在主调函数调用被调函数的前后有一个保存现场和恢复现场的动作。
主调函数保存和恢复现场的通用逻辑如下：

//开辟栈空间，压栈 BP 保存现场
    SUBQ $x, SP    //为函数开辟裁空间
    MOVQ BP, y(SP) //保存当前函数 BP 到 y(SP）位直， y 为相对 SP 的偏移量
    LEAQ y(SP), BP //重直 BP，使其指向刚刚保存 BP 旧值的位置，这里主要
                   //是方便后续 BP 的恢复
//弹出栈，恢复 BP
    MOVQ y(SP), BP //恢复 BP 的值为调用前的值
    ADDQ $x, SP    //恢复 SP 的值为函数开始时的位
*/

/*

	汇编基础

Go 编译器产生的汇编代码是一种中间抽象态，它不是对机器码的映射，而是和平台无关的一个中间态汇编描述，所以汇编代码中有些寄存器是真实的，有些是抽象的，几个抽象的寄存器如下：

SB (Static base pointer)：静态基址寄存器，它和全局符号一起表示全局变量的地址。
FP (Frame pointer)：栈帧寄存器，该寄存器指向当前函数调用栈帧的栈底位置。
PC (Program counter)：程序计数器，存放下一条指令的执行地址，很少直接操作该寄存器，一般是 CALL、RET 等指令隐式的操作。
SP (Stack pointer)：栈顶寄存器，一般在函数调用前由主调函数设置 SP 的值对栈空间进行分配或回收。


Go 汇编器采用 AT&T 风格的汇编
下面代码的分析基于 AMD64 位架构下的 Linux 环境。

*/

func swap(a, b int) (x int, y int) {
	x = b
	y = a
	return
}

/*
编译生成汇编如下:

//- S 产生汇编的代码
//- N 禁用优化
//- 1 禁用内联

GOOS=linux GOARCH=amd64 go tool compile -1 -N -S swap.go >swap.s 2>&1


从汇编的代码得知：
函数的调用者负责环境准备，包括为参数和返回值开辟栈空间。
寄存器的保存和恢复也由调用方负责。
函数调用后回收栈空间，恢复 BP 也由主调函数负责。

*/
