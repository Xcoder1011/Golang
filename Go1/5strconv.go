package main

import (
	"fmt"
	"strconv"
)

// strconv包：字符串和数值类型的相互转换

func main() {

	itoaTest()

	parseTest()

	formatTest()

	appendTest()
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
		fmt.Printf("%v 转换失败！\n", str) //打印： 2 转换失败！
	} else {
		fmt.Printf("tpe: %T, value: %v\n", b1, b1)
	}

	// 2. ParseInt()
	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	// base 指定进制，取值范围是 2 到 36。如果 base 为 0，则会从字符串前置判断，“0x”是 16 进制，“0”是 8 进制，否则是 10 进制。
	// bitSize 指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64。
	// ParseUint() 函数的功能类似于 ParseInt() 函数，但 ParseUint() 函数不接受正负号，用于无符号整型
	str = "-99"
	i1, error := strconv.ParseInt(str, 10, 0)
	if error != nil {
		fmt.Printf("%v 转换失败！\n", i1)
	} else {
		fmt.Printf("tpe: %T, value: %v\n", i1, i1) //打印 tpe: int64, value: -99
	}

	// 3. ParseFloat()
	// func ParseFloat(s string, bitSize int) (f float64, err error)
	// bitSize 指定了返回值的类型，32 表示 float32，64 表示 float64；
	str = "3.1415926"
	f1, error := strconv.ParseFloat(str, 64)
	if error != nil {
		fmt.Printf("%v 转换失败！\n", f1)
	} else {
		fmt.Printf("tpe: %T, value: %v\n", f1, f1) //打印 tpe: float64, value: 3.1415926
	}
}

// Format 系列函数
// Format 系列函数实现了将给定类型数据格式化为字符串类型的功能，
// 其中包括 FormatBool()、FormatInt()、FormatUint()、FormatFloat()。
func formatTest() {

	num1 := true
	str := strconv.FormatBool(num1)
	fmt.Printf("type: %T, value: %v\n ", str, str) // type: string, value: true

	var num2 int64 = 100
	str = strconv.FormatInt(num2, 10)
	fmt.Printf("type: %T, value: %v\n ", str, str) // type: string, value: 100

	// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	// bitSize 表示参数 f 的来源类型（32 表示 float32、64 表示 float64），会据此进行舍入。
	// fmt 表示格式，可以设置为“f”表示 -ddd.dddd、“b”表示 -ddddp±ddd，指数为二进制、
	// “e”表示 -d.dddde±dd 十进制指数、“E”表示 -d.ddddE±dd 十进制指数、“g”表示指数很大时用“e”格式，否则“f”格式、“G”表示指数很大时用“E”格式，否则“f”格式。
	// prec 控制精度（排除指数部分）：当参数 fmt 为“f”、“e”、“E”时，它表示小数点后的数字个数；当参数 fmt 为“g”、“G”时，它控制总的数字个数。
	// 如果 prec 为 -1，则代表使用最少数量的、但又必需的数字来表示 f。
	var f6 float64 = 3.1415926
	str = strconv.FormatFloat(f6, 'E', -1, 64)
	fmt.Printf("type: %T, value: %v\n ", str, str) //  type: string, value: 3.1415926E+00
}

// Append 系列函数
// Append 系列函数用于将指定类型转换成字符串后追加到一个切片中，
// 其中包含 AppendBool()、AppendFloat()、AppendInt()、AppendUint()。
func appendTest() {

	b10 := []byte("int (base10):")
	// func AppendInt(dst []byte, i int64, base int) []byte
	// base ：进制 2 <= base <= 36.
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10)) ///打印： int (base10):-42

	b16 := []byte("int (base16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16)) ///打印： int (base16):-2a
}
