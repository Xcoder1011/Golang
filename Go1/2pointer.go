package main

import (
	"flag"
	"fmt"
)

/*
	指针（pointer）在Go语言中可以被拆分为两个核心概念：
	① 类型指针，允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，而无须拷贝数据，类型指针不能进行偏移和运算。
	② 切片，由指向起始元素的原始指针、元素数量和容量组成。

	Go语言的指针类型变量即拥有指针高效访问的特点，又不会发生指针偏移，从而避免了非法修改关键性数据的问题。
	同时，垃圾回收也比较容易对不会发生偏移的指针进行检索和回收。
	切片比原始指针具备更强大的特性，而且更为安全。切片在发生越界时，运行时会报出宕机，并打出堆栈，而原始指针只会崩溃。


	C/C++中的指针
	指针是 C/C++ 语言拥有极高性能的根本所在，在操作大块数据和做偏移时即方便又便捷
	C/C++ 中指针饱受诟病的根本原因是指针的运算和内存释放
	我们的计算机操作系统经常需要更新、修复漏洞的本质，就是为解决指针越界访问所导致的“缓冲区溢出”的问题。
*/

func main() {

	fmt.Println("Hello World!")

	/*
		一个指针变量可以指向任何一个值的内存地址
		指针变量所指向的值的内存地址在 32 和 64 位机器上分别占用 4 或 8 个字节，占用字节的大小与所指向的值的大小无关。
		当一个指针被定义后没有分配到任何变量时，它的默认值为 nil。指针变量通常缩写为 ptr。
		变量名前面添加&操作符（前缀）来获取变量的内存地址（取地址操作）
	*/
	var cat int = 1
	var str string = "apple"
	fmt.Printf("%p %p", &cat, &str) // 0xc000014168 0xc000010240

	// 格式 ptr := &v    // v 的类型为 T
	// 其中 v 代表被取地址的变量，
	// 变量 v 的地址使用变量 ptr 进行接收，
	// ptr 的类型为*T，称做 T 的指针类型，*代表指针。

	test1()

	test2()

	flagTest()

	test4()
}

// 1. 从指针获取指针指向的值

func test1() {

	var name = "wushangkun"

	// 对字符串取地址, ptr类型为*string
	ptr := &name

	// 打印ptr的类型
	fmt.Printf("\nptr的类型为: %T\n", ptr) // ptr的类型为: *string

	// 打印ptr的指针地址
	fmt.Printf("ptr指针地址： %p\n", ptr) // ptr指针地址： 0xc000010250

	// 对指针进行取值操作
	value := *ptr

	fmt.Printf("取值后的类型: %T\n", value) // 取值后的类型: string

	fmt.Printf("对指针取值后的值： %s\n", value) // 对指针取值后的值： wushangkun

	// 取地址操作符&和取值操作符*是一对互补操作符，
	// &取出地址，*根据地址取出地址指向的值。
}

// 2. 使用指针修改值

func test2() {

	x, y := 1, 2

	// 交换变量值
	swap(&x, &y)

	fmt.Println(x, y) // 2 1
}

// 交换函数
// 定义一个交换函数，参数为 a、b，类型都为 *int 指针类型
func swap(a, b *int) {
	t := *a // 取a指针的值, 赋给临时变量t, t 此时是 int 类型。
	*a = *b // 取b指针的值, 赋给a指针指向的变量
	*b = t  // 将a指针的值赋给b指针指向的变量

	// 总结：
	// 当操作符*在右值时，就是取指向变量的值，
	// 当操作符*在左值时，就是将值设置给指向的变量。
}

// 3. 使用指针变量获取命令行的输入信息

var mode = flag.String("mode", "", "process mode")

func flagTest() {

	flag.Parse()
	fmt.Println(*mode) // fast

	// 命令行运行： go run 2pointer.go --mode=fast
	// 命令行输出结果：fast

	// flag.String 注册了一个名为 mode 的命令行参数，flag 底层知道怎么解析命令行，
	// 并且将值赋给 mode*string 指针，在 Parse 调用完毕后，无须从 flag 获取值，
	// 而是通过自己注册的这个 mode 指针获取到最终的值
}

// 4. 创建指针的另一种方法——new() 函数

func test4() {

	// new() 函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值。
	name := new(string)

	*name = "wushangkun"

	fmt.Println(*name)  // wushangkun
}
