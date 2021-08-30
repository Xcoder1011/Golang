/*

	Test功能测试函数

Go语言自带了 testing 测试包，可以进行自动化的单元测试，输出结果验证，并且可以测试性能。
Go语言的 testing 包提供了三种测试方式，分别是单元（功能）测试、性能（压力）测试和覆盖率测试。


编写测试用例有以下几点需要注意：

测试用例文件不会参与正常源码的编译，不会被包含到可执行文件中；
测试用例的文件名必须以_test.go结尾；
需要使用 import 导入 testing 包；
测试函数的名称要以Test或Benchmark开头，后面可以跟任意字母组成的字符串，但第一个字母必须大写，例如 TestAbc()，一个测试用例文件中可以包含多个测试函数；
单元测试则以(t *testing.T)作为参数，性能测试以(t *testing.B)做为参数；
测试用例文件使用go test命令来执行，源码中不需要 main() 函数作为入口，所有以_test.go结尾的源码文件内以Test开头的函数都会自动执行。

*/

package demo

import "testing"

/// 1）单元（功能）测试

// 命令： go test -v
// 要开始一个单元测试，需要准备一个 go 源码文件，在命名文件时文件名必须以_test.go结尾，
// 每个测试用例的名称需要以 Test 为前缀， 例如：

// func TestXxx( t *testing.T ){
//     //......
// }

func TestGetArea(t *testing.T) { // 单元测试则以(t *testing.T)作为参数
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	}

	/*
		执行测试命令，运行结果如下所示：

		bash-3.2$ go test -v
		=== RUN   TestGetArea
		--- PASS: TestGetArea (0.00s)
		PASS
		ok
		bash-3.2$

	*/
}

/// 2）性能（压力）测试
//  命令：go test -bench='.'
func BenchmarkGetArea(t *testing.B) { // 性能测试以(t *testing.B)做为参数
	for i := 0; i < t.N; i++ {
		GetArea(40, 50)
	}

	/*

		bash-3.2$ go test -bench='.'
		goos: darwin
		goarch: amd64
		pkg: Test
		cpu: Intel(R) Core(TM) i5-7500 CPU @ 3.40GHz
		BenchmarkGetArea-4   	1000000000 		 0.2755 ns/op
		PASS
		ok
		bash-3.2$

		上面信息显示了程序执行 1000000000 次，共耗时 0.2755 纳秒。
	*/
}

/// 3）覆盖率测试
//  命令：go test -cover
// 覆盖率测试能知道测试程序总共覆盖了多少业务代码（也就是 demo_test.go 中测试了多少 demo.go 中的代码），可以的话最好是覆盖100%。

/*

	bash-3.2$ go test -cover
	PASS
	coverage: 100.0% of statements
	ok
	bash-3.2$


	// 覆盖率:100%。
*/
