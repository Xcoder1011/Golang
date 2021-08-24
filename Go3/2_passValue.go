package main

import (
	"fmt"
)

// 用于测试值传递效果的结构体
type Data struct {
	complax []int // 测试切片在参数传递中的效果

	instance InnerData // 实例分配的innerData

	ptr *InnerData // 将ptr声明为InnerData的指针类型
}

type InnerData struct {
	a int
}

func main() {

	data := Data{
		complax:  []int{2, 3, 4},
		instance: InnerData{3},
		ptr:      &InnerData{1},
	}

	// 输入结构的成员情况
	fmt.Printf("data value: %+v\n", data)
	// 输入结构的指针地址
	fmt.Printf("data ptr: %p\n", &data)
	// 传入结构体，返回同类型的结构体
	out := passByValue(data)
	// 输出结构的成员情况
	fmt.Printf("out value: %+v\n", out)
	// 输出结构的指针地址
	fmt.Printf("out ptr: %p\n", &out)
	
	
	/*
	
	data value: {complax:[2 3 4] instance:{a:3} ptr:0xc0000ae008}
data ptr: 0xc000098180
inFunc value: {complax:[2 3 4] instance:{a:3} ptr:0xc0000ae008}
inFunc ptr: 0xc000098210
out value: {complax:[2 3 4] instance:{a:3} ptr:0xc0000ae008}
out ptr: 0xc0000981e0
	*/

}

func passByValue(inFunc Data) Data {

	// 输出参数的成员情况
	fmt.Printf("inFunc value: %+v\n", inFunc)
	// 打印inFunc的指针
	fmt.Printf("inFunc ptr: %p\n", &inFunc)

	return inFunc
}
