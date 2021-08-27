/*

	匿名函数

	func(参数列表)(返回参数列表){
		函数体
	}
*/

package main

import (
	"flag"
	"fmt"
)

func main() {

	// 1) 在定义时调用匿名函数
	func(num int) {
		fmt.Println("num is ", num) // num is  100
	}(100) // 匿名函数可以在声明后调用

	// 2) 将匿名函数赋值给变量
	f3 := func(num int) {
		fmt.Println("num is ", num) // num is  200
	}
	// 使用f3()调用
	f3(200)

	// 3) 匿名函数用作回调函数
	list := []int{1, 2, 3, 4, 5}
	visit(list, func(v int) {
		fmt.Println(v + 10)
	})

	// 4) 使用匿名函数实现操作封装
	testAnoymous()

}

// 匿名函数用作回调函数

func visit(list []int, f func(int)) {

	for _, v := range list {
		f(v)
	}
}

// 使用匿名函数实现操作封装

var skillParam = flag.String("skill", "run", "skill to perform")

// 从命令行输入 --skill 可以将=后的字符串传入 skillParam 指针变量
// 终端输入： go run 4_anonymous_function.go --skill=fly
// 打印： angel fly
func testAnoymous() {

	// 解析命令行参数，解析完成后，skillParam 指针变量将指向命令行传入的值。
	flag.Parse()

	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	// skillParam 是一个 *string 类型的指针变量
	// 使用 *skillParam 获取到命令行传过来的值
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}
