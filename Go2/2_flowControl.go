package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ { //左花括号{必须与 for 处于同一行。

		sum += i
	}
	fmt.Println(sum) // 45

	// 3) 只有一个循环条件的循环

	var i int
	for i <= 10 {
		i++
		fmt.Println(i)

	}

	sum2 := 0
	for { // 无限循环
		sum2++
		if sum2 > 100 {
			break
		}
	}
	fmt.Println(sum2) // 101

	// Go语言的 for 循环同样支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break，可以选择中断哪一个循环，如下例：

	// for j := 0; j < 5; j++ {
	// 	for i := 0; i < 10; i++ {
	// 		if i > 5 {
	// 			break JLoop
	// 		}
	// 		fmt.Println(i)
	// 	}
	// }
	// JLoop:
	// fmt.Println("break JLoop")

	/*
			for range 可以遍历数组、切片、字符串、map 及通道（channel），
			for range 语法上类似于其它语言中的 foreach 语句，
			一般形式为：
				for key, val := range coll {
		    	...
				}

		需要要注意的是，val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，
		对它所做的任何修改都不会影响到集合中原有的值。
	*/

	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d  value:%d\n", key, value)
	}

	/// 遍历字符串

	var str = "hello 你好"
	for key, value := range str {

		fmt.Printf("key:%d value:0x%x\n", key, value)
		/*
			key:0 value:0x68
			key:1 value:0x65
			key:2 value:0x6c
			key:3 value:0x6c
			key:4 value:0x6f
			key:5 value:0x20
			key:6 value:0x4f60
			key:9 value:0x597d

			代码中的变量 value，实际类型是 rune 类型，以十六进制打印出来就是字符的编码。

		*/
	}

	/// 遍历通道（channel）——接收通道数据

	channel := make(chan int) // 创建一个整型类型的通道。

	go func() { // 启动一个 goroutine，往通道中推送数据 1、2、3，然后结束并关闭通道
		channel <- 1
		channel <- 2
		channel <- 3
		close(channel)
	}()

	for v := range channel {
		fmt.Println(v)
	}

	haskell()

	switchTest()

	gotoTest()

	breakTest()
}

// 输出九九乘法表
func haskell() {
	// 遍历, 决定处理第几行
	for y := 1; y <= 9; y++ {
		// 遍历, 决定这一行有多少列
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		// 手动生成回车
		fmt.Println()
	}

	/*
	   1*1=1
	   1*2=2 2*2=4
	   1*3=3 2*3=6 3*3=9
	   1*4=4 2*4=8 3*4=12 4*4=16
	   1*5=5 2*5=10 3*5=15 4*5=20 5*5=25
	   1*6=6 2*6=12 3*6=18 4*6=24 5*6=30 6*6=36
	   1*7=7 2*7=14 3*7=21 4*7=28 5*7=35 6*7=42 7*7=49
	   1*8=8 2*8=16 3*8=24 4*8=32 5*8=40 6*8=48 7*8=56 8*8=64
	   1*9=9 2*9=18 3*9=27 4*9=36 5*9=45 6*9=54 7*9=63 8*9=72 9*9=81
	*/
}

func switchTest() {

	// switch表达式不需要为常量，甚至不需要为整数
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}

	// 1) 一分支多值, 多个 case 要放在一起

	a = "mum"
	switch a {
	case "mum", "daddy":
		fmt.Println("family")
	}

	// case 后不仅仅只是常量
	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}

	// 跨越 case 的 fallthrough——兼容C语言的 case 设计

	// 在Go语言中 case 是一个独立的代码块，执行完毕后不会像C语言那样紧接着执行下一个 case，
	// 但是为了兼容一些移植代码，依然加入了 fallthrough 关键字来实现这一功能，
	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}
	/*
		输出:
		hello
		world
	*/
}

func gotoTest() {

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}

	// 手动返回, 避免执行进入标签
	return

	// 标签
breakHere:
	fmt.Println("done")
}

func breakTest() {
	// 输出
	// 0 2
	// 1 2
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue OuterLoop // 结束当前循环，开启下一次的外层循环
			case 3:
				fmt.Println(i, j) // 退出 OuterLoop 对应的循环之外
				break OuterLoop
			}
		}
	}
}
