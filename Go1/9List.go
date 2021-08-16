package main

import (
	"container/list"
	"fmt"
)

// 列表是一种非连续的存储容器，由多个节点组成，节点通过一些变量记录彼此之间的关系，列表有多种实现方法，如单链表、双链表等。

// 在Go语言中，列表使用 container/list 包来实现，内部的实现原理是双链表，列表能够高效地进行任意位置的元素插入和删除操作。

func main() {

	// 1. 初始化列表
	// 分别是使用 New() 函数和 var 关键字声明，两种方法的初始化效果都是一致的。

	// 1) 通过 container/list 包的 New() 函数初始化 list
	// 变量名 := list.New()
	l1 := list.New()

	// 2) 通过 var 关键字声明初始化 list
	// var 变量名 list.List
	var l2 list.List

	// 双链表支持从队列前方或后方插入元素，分别对应的方法是 PushFront 和 PushBack。
	// 这两个方法都会返回一个 *list.Element 结构，如果在以后的使用中需要删除插入的元素，
	// 则只能通过 *list.Element 配合 Remove() 方法进行删除，这种方法可以让删除更加效率化，同时也是双链表特性之一。

	l1.PushBack("first")
	l1.PushFront(67)

	element := l2.PushBack("first")
	fmt.Println(element) // &{0xc000100060 0xc000100060 0xc000100060 first}

	removeElement()

	forEach()

}

/// 从列表中删除元素

func removeElement() {
	l := list.New()

	// 尾部添加
	l.PushBack("canon") // canon

	// 头部添加
	l.PushFront(67) // 67, canon

	// 尾部添加后保存元素句柄
	element := l.PushBack("first") // 67, canon, first

	// 在fist之后添加high
	l.InsertAfter("high", element) // 67, canon, first, high

	// 在fist之前添加noon
	l.InsertBefore("noon", element) // 67, canon, noon, first, high

	// 移除 element 变量对应的元素。
	l.Remove(element) // 67, canon, noon, high

	fmt.Println(l) // &{{0xc0001001b0 0xc000100210 <nil> <nil>} 4}

}

func forEach() {

	l := list.New()

	// 尾部添加
	l.PushBack("canon") // canon

	// 头部添加
	l.PushFront(67) // 67, canon

	// i:=l.Front() 表示初始赋值
	// next , preview
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
		/*

			67
			canon

		*/
	}
}
