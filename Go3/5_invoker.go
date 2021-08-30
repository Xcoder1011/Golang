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
// 函数的声明不能直接实现接口
// 需要将函数定义为类型后，使用类型实现结构体
type FuncCaller func(interface{}) // 将 func(interface{}) 定义为 FuncCaller 类型。

// 函数体实现Invoker的Call
// 当类型方法被调用时，还需要调用函数本体。
func (f FuncCaller) Call(p interface{}) { // FuncCaller 的 Call() 方法将实现 Invoker 的 Call() 方法。

	// FuncCaller 的 Call() 方法被调用与 func(interface{}) 无关，还需要手动调用函数本体。
	f(p) // 调用f函数本体
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
	invoker = s // s 类型为 *Struct，已经实现了 Invoker 接口类型，因此赋值给 invoker 时是成功的。

	// 使用接口调用实例化结构体的方法 Struct.Call
	invoker.Call("hello, wushangkun") // 打印： from struct :  hello, wushangkun

	// ------------ 例2：函数实现接口

	// 将匿名函数转为 FuncCaller类型， 再赋值给接口
	// 函数来源可以是命名函数、匿名函数或闭包
	invoker = FuncCaller(func(v interface{}) { // FuncCaller 无须被实例化，只需要将函数转换为 FuncCaller 类型即可
		fmt.Println("from function: ", v)
	})

	// 使用接口调用FuncCaller.Call， 内部会调用函数本体
	invoker.Call("hello, go") // 打印： from function:  hello, go
}

/*

//// -------------- HTTP包中的例子

// Handler接口 用于定义每个 HTTP 的请求和响应的处理过程。
type Handler interface {
	ServerHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

// HandlerFunc 类型实现了 Handler 的 ServeHTTP 方法
func (f HandlerFunc) ServerHTTP(w ResponseWriter, r *Request) {
	// 底层可以同时使用各种类型来实现 Handler 接口进行处理。
	f(w, r)
}

// 使用闭包实现默认的 HTTP 请求处理，可以使用 http.HandleFunc() 函数
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

// DefaultServeMux 是 ServeMux 结构，拥有 HandleFunc() 方法
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {

	// 将外部传入的函数 handler() 转为 HandlerFunc 类型
	mux.Handle(pattern, HandleFunc(handler))
}

*/
