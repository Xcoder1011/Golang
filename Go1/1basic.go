package main

import (
	"encoding/base64"
	"fmt"
	"strings"
	"unicode/utf8"
)

// 声明全局变量
// 全局变量声明必须以 var 关键字开头，如果想要在外部包中使用全局变量的首字母必须大写。
var globalName string
var globalAge int = 20

// 常量
const cacheKey string = "cacheKey" // 相当于 math.Pi 的近似值
const IPv4Len = 4

// 批量声明多个常量
const (
	e  = 2.7182818
	pi = 3.1415926
)

const (
	a = 1
	b
	c = 2
	d
)

// fmt.Println(a, b, c, d) // "1 1 2 2"

// Go语言现阶段没有枚举类型，可以利用iota 来模拟枚举类型
type Weekday int

const (
	Sunday  Weekday = iota // 从周日 0 开始， 枚举值自增1
	Monday                 // 1
	Tuesday                // 2
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	FlagNone = 1 << iota // 移位操作
	FlagRed
	FlagGreen
	FlagBlue
)

// fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)   // 二进制位左移一位，打印结果： 2 4 8
// fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue)   // 二进制格式输出： 10 100 1000

func main() {

	fmt.Println("Hello World!")

	/*
		Go语言的基本类型有：
		bool
		string
		int、int8、int16、int32、int64
		uint、uint8、uint16、uint32、uint64、uintptr
		byte // uint8 的别名
		rune // int32 的别名 代表一个 Unicode 码
		float32、float64
		complex64、complex128
	*/

	// // 申明局部变量
	// var a string
	// var b int
	// var c float32
	// var d *int
	// var e,f byte  // uint8 的别名

	/*
		当一个变量被声明之后，系统自动赋予它该类型的零值：
		int 为 0，float 为 0.0，bool 为 false，string 为空字符串，布尔型变量默认为 bool, 切片、函数、指针变量的默认为 nil 等。
		所有的内存在 Go 中都是经过初始化的。

		骆驼命名法

	*/

	// 申明局部变量hp
	// Go语言程序中全局变量与局部变量名称可以相同，但是函数体内的局部变量会被优先考虑。
	// var hp int = 100 // 等同于下面的
	// var hp = 100  // 编译器推导类型
	// hp := 100     // 短变量声明精简的写法

	// 批量格式
	// var (
	// 	a1 int
	// 	a2 []float32
	// 	a3 func() bool

	// 	a4 struct {
	// 		x int
	// 	}
	// )

	// 变量交换 方式1
	var c1 = 100
	var c2 = 90

	c1 = c1 ^ c2
	c2 = c2 ^ c1
	c1 = c1 ^ c2

	fmt.Println(c1, c2)

	// 变量交换 方式2 （Go 多重赋值）
	var c3 = 100
	var c4 = 90

	c4, c3 = c3, c4 // 多重赋值时，变量的左值和右值按从左到右的顺序赋值。

	fmt.Println(c3, c4)

	// float32 类型的浮点数可以提供大约 6 个十进制数的精度，而 float64 则可以提供约 15 个十进制数的精度，
	// 通常应该优先使用 float64 类型，因为 float32 类型的累计计算误差很容易扩散，
	// 并且 float32 能精确表示的正整数并不是很大。

	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!

	// 复数
	// 是由两个浮点数表示的，其中一个表示实部（real），一个表示虚部（imag）。
	// x、y 分别表示构成该复数的两个 float64 类型的数值，x 为实部，y 为虚部。

	/*var name complex128 = complex(x, y)

	// 对于一个复数z := complex(x, y)，可以通过Go语言的内置函数real(z) 来获得该复数的实部，也就是 x；
	// 通过imag(z) 获得该复数的虚部，也就是 y。
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // "(-5+10i)"
	fmt.Println(real(x * y))         // "-5"
	fmt.Println(imag(x * y))    */ // "10"

	// 字符串拼接符“+”
	s := "hel" + "lo,"
	s += "world!"
	fmt.Println(s) //输出 “hello, world!”

	// 双引号: 字符串字面量（string literal）
	// 多行字符串时，就必须使用`反引号
	const str = `第一行
	第二行
	第三行
	\r\n
	`
	fmt.Println(str)

	//  len()可以用来获取切片、字符串、通道（channel）等的长度
	tip1 := "genji is a ninja"
	fmt.Println(len(tip1))

	// Unicode 字符串长度使用 utf8.RuneCountInString() 函数。
	fmt.Println(utf8.RuneCountInString("忍者"))

	// 遍历每一个ASCII字符
	theme := "狙击 start"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii: %c  %d\n", theme[i], theme[i])
	}

	// 字符串截取
	// strings.Index：正向搜索子字符串。
	// strings.LastIndex：反向搜索子字符串。
	tracer := "死神来了, 死神bye bye"
	comma := strings.Index(tracer, ", ")
	fmt.Println(comma)

	var progress = 2
	var target = 8
	// 两参数格式化
	title := fmt.Sprintf("已采集%d个药草, 还需要%d个完成任务", progress, target)
	fmt.Println(title)

	pi := 3.14159
	// 按数值本身的格式输出
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true) // "月球基地" 3.14159 true
	fmt.Println(variant)

	// 匿名结构体声明, 并赋予初值
	profile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   150,
	}
	fmt.Printf("使用'%%+v' %+v\n", profile)
	fmt.Printf("使用'%%#v' %#v\n", profile)
	fmt.Printf("使用'%%T' %T\n", profile)

	/*
		已采集2个药草, 还需要8个完成任务
		"月球基地" 3.14159 true
		使用'%+v' &{Name:rat HP:150}
		使用'%#v' &struct { Name string; HP int }{Name:"rat", HP:150}
		使用'%T' *struct { Name string; HP int }

		动  词	功  能
		%v	按值的本来值输出
		%+v	在 %v 基础上，对结构体字段名和值进行展开
		%#v	输出 Go 语言语法格式的值
		%T	输出 Go 语言语法格式的类型和值
		%%	输出 % 本体
		%b	整型以二进制方式显示
		%o	整型以八进制方式显示
		%d	整型以十进制方式显示
		%x	整型以十六进制方式显示
		%X	整型以十六进制、字母大写方式显示
		%U	Unicode 字符
		%f	浮点数
		%p	指针，十六进制方式显示

	*/

	// Base64 编码
	message := "https://golang.org/"
	encodeMessage := base64.StdEncoding.EncodeToString([]byte(message)) // 字符串需要转换为字节数组
	fmt.Println(encodeMessage)                                          // aHR0cHM6Ly9nb2xhbmcub3JnLw==

	// Base64 解码
	data, err := base64.StdEncoding.DecodeString(encodeMessage)
	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Println(string(data)) // 将返回的字节数组（[]byte）转换为字符串。
	}

	// 字符

	// 一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
	// 另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。

	// var ch1 byte = 'A'
	// var ch2 byte = 65     // 在 ASCII 码表中，A 的值是 65
	// var ch3 byte = '\x41' // 使用 16 进制表示则为 41

	// 书写 Unicode 字符时，需要在 16 进制数之前加上前缀\u或者\U
	var ch int = '\u0041'
	var ch2 int = '\u03B2'                     // 如果需要使用到 4 字节，则使用\u前缀
	var ch3 int = '\U00101234'                 // 如果需要使用到 8 个字节，则使用\U前缀
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point

	/*
		输出：
		65 - 946 - 1053236
		A - β - r
		41 - 3B2 - 101234
		U+0041 - U+03B2 - U+101234
	*/

	/*
		nil 标识符是不能比较的
		输出  invalid operation: nil == nil (operator == not defined on nil)
		// fmt.Println(nil == nil)

		nil 不是关键字或保留字
		我们可以定义一个名称为 nil 的变量
		var nil = errors.New("my god") //  use of untyped nil
		fmt.Printf("%T", nil)

		nil 是 map、slice、pointer、channel、func、interface 的零值
		不同类型的 nil 值占用的内存大小可能是不一样的

		var p *struct{}
		fmt.Println(unsafe.Sizeof(p)) // 8
		var s []int
		fmt.Println(unsafe.Sizeof(s)) // 24
		var m map[int]bool
		fmt.Println(unsafe.Sizeof(m)) // 8
		var c chan string
		fmt.Println(unsafe.Sizeof(c)) // 8
		var f func()
		fmt.Println(unsafe.Sizeof(f)) // 8
		var i interface{}
		fmt.Println(unsafe.Sizeof(i)) // 16

	*/

}

// 形式参数
// 在函数未被调用时，函数的形参并不占用实际的存储单元，也没有实际值。
func sum(a, b int) int {
	num := a + b
	return num
}

// 匿名变量

// 下画线“_”，“_”本身就是一个特殊的标识符，被称为空白标识符。
// 匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。
// 它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），
// 但任何赋给这个标识符的值都将被抛弃，因此这些值不能在后续的代码中使用，

func getData() (int, int) {
	return 100, 200
}

func test3() {
	d1, _ := getData() // 只需要获取第一个返回值，所以将第二个返回值的变量设为下画线（匿名变量）。
	_, d2 := getData() // 将第一个返回值的变量设为匿名变量。
	fmt.Println(d1, d2)
}

// 多重赋值

type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}

func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// 简短格式
// 名字 := 表达式

// 定义变量，同时显式初始化。
// 不能提供数据类型。
// 只能用在函数内部。

func test() {
	b1 := 10
	b1, b2 := 1, 2
	b3, b4 := 3, "abc"
	fmt.Printf("%d - %d - %d - %v\n", b1, b2, b3, b4) // integer
}
