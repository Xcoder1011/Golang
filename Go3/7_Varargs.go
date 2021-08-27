package main

import (
	"bytes"
	"fmt"
)

/*
	可变参数（变参函数）

*/
func main() {

	/// 1. 可变参数类型

	// ...type格式的类型只能作为函数的参数类型存在，并且必须是最后一个参数，
	// 它是一个语法糖（syntactic sugar），即这种语法对语言的功能并没有影响，

	myfunc(2, 3, 4)
	myfunc(1, 3, 7, 13)

	/// 2. 任意类型的可变参数

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)

	/// 3. 遍历可变参数列表

	// 输入3个字符串, 将它们连成一个字符串
	fmt.Println(joinStrings("pig ", "and", " rat"))             // pig and rat
	fmt.Println(joinStrings("hammer", " mom", " and", " hawk")) // hammer mom and hawk

	// 获得每一个参数的类型
	fmt.Println(printTypeValue(100, "str", true))

	/*

		value: 100 type: int
		value: str type: string
		value: true type: bool
	*/

	/// 4. 在多个可变参数函数中传递参数
	print(1, 2, 3)

}

// 可变参数类型

// myfunc() 接受不定数量的参数，这些参数的类型全部是 int
// 从内部实现机理上来说，类型...type本质上是一个数组切片，也就是[]type，
// 这也是为什么参数 args 可以用 for 循环来获得每个传入的参数。
func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

// 任意类型的可变参数

// 如果你希望传任意类型，可以指定类型为 interface{}，
// 下面是Go语言标准库中 fmt.Printf() 的函数原型：
func Printf(format string, args ...interface{}) {
	// ...
	// 用 interface{} 传递任意类型数据是Go语言的惯例用法
	// 使用 interface{} 仍然是类型安全的，这和 C/C++ 不太一样
}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

// 遍历可变参数列表
// 定义一个函数, 参数数量为0~n, 类型约束为字符串
func joinStrings(slist ...string) string {
	// 定义一个字节缓冲, 快速地连接字符串
	var b bytes.Buffer
	// 遍历可变参数列表slist, 类型为[]string
	for _, s := range slist {
		// 将遍历出的字符串连续写入字节数组
		b.WriteString(s)
	}
	// 将连接好的字节数组转换为字符串并输出
	return b.String()
}

/// 获得可变参数类型  ——--  获得每一个参数的类型

func printTypeValue(slist ...interface{}) string {
	// 字节缓冲作为快速字符串连接
	var b bytes.Buffer
	// 遍历参数
	for _, s := range slist {
		// 将interface{}类型格式化为字符串
		// 使用 fmt.Sprintf 配合%v动词，可以将 interface{} 格式的任意值转为字符串。
		str := fmt.Sprintf("%v", s)
		// 类型的字符串描述
		var typeString string
		// 对s进行类型断言
		switch s.(type) {
		case bool: // 当s为布尔类型时
			typeString = "bool"
		case string: // 当s为字符串类型时
			typeString = "string"
		case int: // 当s为整型类型时
			typeString = "int"
		}
		// 写字符串前缀
		b.WriteString("value: ")
		// 写入值
		b.WriteString(str)
		// 写类型前缀
		b.WriteString(" type: ")
		// 写类型字符串
		b.WriteString(typeString)
		// 写入换行符
		b.WriteString("\n")
	}
	return b.String()
}

/// 在多个可变参数函数中传递参数

// 实际打印的函数
func rawPrint(rawList ...interface{}) {
	// 遍历可变参数切片
	for _, a := range rawList {
		// 打印参数
		fmt.Println(a)
	}
}

// 打印函数封装
func print(slist ...interface{}) {
	// 将slist可变参数切片完整传递给下一个函数
	// 可以在传递时给可变参数变量后面添加...，这样就可以将切片中的元素进行传递
	rawPrint(slist...)
}
