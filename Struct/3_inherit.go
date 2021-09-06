package main

import (
	"fmt"
)

/*
	“继承”特性

	Go语言结构体内嵌模拟类的继承
	Go语言的结构体内嵌特性就是一种组合特性，使用组合特性可以快速构建对象的不同特性。

*/

// 可飞行的
type Flying struct {
	condition int
}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

// 可行走的
type Walkable struct {
	condition int
}

func (w *Walkable) Walk() {
	fmt.Println("can walk")
}

// 人类
type Human struct {
	Walkable // 人类能行走
}

type Bird struct {
	Flying   // 鸟类能飞行
	Walkable // 鸟类能行走
}

// 内嵌匿名结构体
type Dog struct {
	Walkable          // 能行走
	Jump     struct { // 能跳
		condition int
	}
}

func main() {

	b := new(Bird)
	b.Fly()
	b.Walk()

	h := new(Human)
	h.Walk()

	// 初始化内嵌结构体
	bird := Bird{
		Flying: Flying{
			condition: 1,
		},
		Walkable: Walkable{
			condition: 0,
		},
	}

	fmt.Printf("%+v\n", bird) // {Flying:{condition:1} Walkable:{condition:0}}

	// 初始化内嵌结构体
	dog := Dog{
		Walkable: Walkable{
			condition: 1,
		},
		Jump: struct {
			condition int
		}{
			condition: 1,
		},
	}

	fmt.Printf("%+v\n", dog) // {Flying:{condition:1} Walkable:{condition:0}}

}
