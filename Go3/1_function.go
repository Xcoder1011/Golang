/*

函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。

func 函数名(形式参数列表)(返回值列表){
    函数体
}
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hypot(3, 4))  // "5"
	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"

	// Go语言支持多返回值
	// Go语言既支持安全指针，也支持多返回值
	// conn, err := connectToNetwork()

	// 1) 同一种类型返回值
	a, b := typedTwoValues()
	fmt.Println(a, b) // 1 2

	// 2) 带有变量名的返回值
	c, d := namedRetValues()
	fmt.Println(c, d) // 1 2

	// 将返回值作为打印参数
	fmt.Println(resolveTime(1000)) // 0 0 16

	// 只获取消息和分钟
	_, hour, minute := resolveTime(18000)
	fmt.Println(hour, minute) // 5 300

	// 只获取天
	day, _, _ := resolveTime(90000)
	fmt.Println(day) // 1

	/// 函数变量  --- 把函数作为值保存到变量中
	var f func()
	f = testFunction
	f()

}

func testFunction() {
	fmt.Println("testFunction")
}

/*
	func 函数名(形式参数列表)(返回值列表){
    		函数体
	}
*/

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// 如果一组形参或返回值有相同的类型，我们不必为每个形参都写出参数类型，下面 2 个声明是等价的：
func f1(i, j, k int, s, t string)                { /* ... */ }
func f2(i int, j int, k int, s string, t string) { /* ... */ }

func add(x int, y int) int { return x + y }
func sub(x, y int) (z int) { z = x - y; return }

func first(x int, _ int) int { return x } // 空白标识符_可以强调某个参数未被使用。
func zero(int, int) int      { return 0 }

// 在函数中，实参通过值传递的方式进行传递，因此函数的形参是实参的拷贝，对形参进行修改不会影响实参，
// 但是，如果实参包括引用类型，如指针、slice(切片)、map、function、channel 等类型，实参可能会由于函数的间接引用被修改。

// 1) 同一种类型返回值
func typedTwoValues() (int, int) {
	return 1, 2
}

// 2) 带有变量名的返回值
func namedRetValues() (a, b int) {
	a = 1
	b = 2
	return
}

// 等同于
func namedRetValues2() (a, b int) {
	a = 1
	return a, 2
}

const (
	// 定义每分钟的秒数
	SecondsPerMinute = 60
	// 定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	// 定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

func resolveTime(seconds int) (day, hour, minute int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute
	return
}
