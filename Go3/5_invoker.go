package main

import (
	"fmt"
)

/*
	函数类型实现接口 ——-- 把函数作为接口来调用
*/

// 调用器接口
type Invoker interface {
	// 这个接口需要实现 Call() 方法，
	// 调用时会传入一个 interface{} 类型的变量，这种类型的变量表示任意类型的值。
	Call(interface{})
}

// 函数定义为类型
type FuncCaller func(interface{})

// 函数体实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {

	// 调用f函数本体
	f(p)
}

// 结构体类型
// 定义结构体
type Struct struct {
	name string
}

// 结构体实现Invoker的Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct : ", p)
}

func main() {

	// 声明接口变量
	var invoker Invoker

	// ------------ 例1：结构体实现接口

	// 实例化结构体
	// new函数的参数是一个类型，并且返回一个指向该类型内存地址的指针, 即类型 *Struct
	var s *Struct = new(Struct)
	s.name = "test"

	// 将实例化的结构体赋值到接口
	invoker = s

	// 使用接口调用实例化结构体的方法 Struct.Call
	invoker.Call("hello, wushangkun") // 打印： from struct :  hello, wushangkun

	// ------------ 例2：函数实现接口

	// 将匿名函数转为 FuncCaller类型， 再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function: ", v)
	})

	// 使用接口调用FuncCaller.Call， 内部会调用函数本体
	invoker.Call("hello, go") // 打印： from function:  hello, go
}
