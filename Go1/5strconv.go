package main

import (
	"fmt"
	"strconv"
)

// strconv包：字符串和数值类型的相互转换

func main() {

	itoaTest()

	parseTest()

}

// 整型 ~ 字符串
func itoaTest() {

	// 1.整型转字符串
	num := 100
	str := strconv.Itoa(num)
	fmt.Printf("tpe: %T, value: %v\n", str, str) //打印： tpe: string, value: 100

	// 2.字符串转整型
	str = "s100"
	num2, error := strconv.Atoi(str)
	if error != nil {
		fmt.Printf("%v 转换失败！\n", str) //打印： s100 转换失败！
	} else {
		fmt.Printf("tpe: %T, value: %v\n", num2, num2)
	}

	str = "100"
	num3, error := strconv.Atoi(str)
	if error != nil {
		fmt.Printf("%v 转换失败！\n", str)
	} else {
		fmt.Printf("tpe: %T, value: %v\n", num3, num3) //打印： tpe: int, value: 100
	}
}

// Parse 系列函数
// Parse 系列函数用于将字符串转换为指定类型的值，其中包括 ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
func parseTest() {

	// 1. ParseBool()
	// ParseBool() 函数用于将字符串转换为 bool 类型的值，
	// 它只能接受 1、0、t、f、T、F、true、false、True、False、TRUE、FALSE，其它的值均返回错误，

	str := "1"
	b, error := strconv.ParseBool(str)
	if error != nil {
		fmt.Printf("%v 转换失败！\n", str)
	} else {
		fmt.Printf("tpe: %T, value: %v\n", b, b) //打印： tpe: bool, value: true
	}

	str = "2"
	b1, error := strconv.ParseBool(str)
	if error != nil {
		fmt.Printf("%v 转换失败！\n", str)
	} else {
		fmt.Printf("tpe: %T, value: %v\n", b1, b1) //打印： tpe: bool, value: true
	}

}
