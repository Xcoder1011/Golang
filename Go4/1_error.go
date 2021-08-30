package main

import (
	"errors"
	"fmt"
)

/*
	处理运行时错误

Go语言的错误处理思想及设计包含以下特征：
一个可能造成错误的函数，需要返回值中返回一个错误接口（error），如果调用是成功的，错误接口将返回 nil，否则返回错误。
在函数调用后需要检查错误，如果发生错误，则进行必要的错误处理。

*/

type error interface {
	Error() string
}

func main() {

	// 在Go语言中，使用 errors 包进行错误的定义
	var err = errors.New("this is an error")
	fmt.Println(err) // this is an error

	/// 2) 在代码中使用错误定义

	fmt.Println(div(1, 0)) // 0 division by zero

	/// 3) 在解析中使用自定义错误

	var e error
	// 创建一个错误实例，包含文件名和行号
	e = newParseError("main.go", 1)

	// 通过error接口查看错误描述
	// 错误对象都要实现 error 接口的 Error() 方法返回错误描述，这样，所有的错误都可以获得字符串的描述
	fmt.Println(e.Error()) // main.go:1

	// 根据错误接口具体的类型，获取详细错误信息
	switch detail := e.(type) {
	case *ParseError: // 这是一个解析错误
		fmt.Printf("Filename: %s Line: %d\n", detail.Filename, detail.Line) // Filename: main.go Line: 1
	default: // 其他类型的错误
		fmt.Println("other error")
	}

}

/// 1) 自定义一个错误

// 错误字符串
type errorString struct {
	s string
}

// 返回发生何种错误
// 实现 error 接口的 Error() 方法
func (e *errorString) Error() string {
	return e.s
}

// 创建错误对象
func New(text string) error {
	return &errorString{text}
}

/// 2) 在代码中使用错误定义

// 定义除数为0的错误
var errDivisionByZero = errors.New("division by zero")

func div(dividend, divisor int) (int, error) {
	// 判断除数为0的情况并返回
	if divisor == 0 {
		return 0, errDivisionByZero
	}
	// 正常计算，返回空错误
	return dividend / divisor, nil
}

/// 3) 在解析中使用自定义错误

// 声明一个解析错误
type ParseError struct {
	Filename string // 文件名
	Line     int    // 行号
}

// 实现error接口
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

// 创建一些解析错误
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}
