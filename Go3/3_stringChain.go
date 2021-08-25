package main

import (
	"fmt"
	"strings"
)

func main() {

	// 字符串包含 go 前缀及空格
	strings := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	// 处理函数链
	chain := []func(string) string{
		removePrefix,
		toUpper,
		trimSpace,
	}

	// 处理字符串
	stringProcess(strings, chain)

	for _, str := range strings {
		fmt.Println(str)
	}
	/*
		输出：

			SCANNER
			PARSER
			COMPILER
			PRINTER
			FORMATER
	*/

}

// 字符串处理函数, 传入字符串切片和处理链

// 字符串切片（list[]string）
// 链式处理函数的切片(chain []func(string) string:)
func stringProcess(list []string, chain []func(string) string) {

	// 遍历每一个字符串
	for index, str := range list {
		result := str
		// 遍历每一个处理链
		for _, process := range chain {
			// 输入一个字符串进行处理, 返回数据作为下一个处理链的输入
			result = process(result)
		}
		// 将结果放回切片
		list[index] = result
	}
}

func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func toUpper(str string) string {
	return strings.ToUpper(str)
}

func trimSpace(str string) string {
	return strings.TrimSpace(str)
}
