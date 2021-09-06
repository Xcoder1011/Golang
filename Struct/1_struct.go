package main

import (
	"fmt"
)

/*
	结构体（struct）

type 类型名 struct {
    字段1 字段1类型
    字段2 字段2类型
    …
}

Go 语言中的类型可以被实例化，使用new或&构造的类型实例的类型是类型的指针。

关于 Go 语言的类（class）:
Go 语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。
Go 语言的结构体与“类”都是复合结构体，但 Go 语言中结构体的内嵌配合接口比面向对象具有更高的扩展性和灵活性。
Go 语言不仅认为结构体能拥有方法，且每种自定义类型也可以拥有自己的方法。

*/

// 结构体中的字段名必须唯一。
type Point struct {
	X int
	Y int
}

// 同类型的变量也可以写在一行
type Color struct {
	R, G, B byte
}

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
	int
	Point
}

type Command struct {
	Name    string // 指令名称
	Var     *int   // 指令绑定的变量
	Comment string // 指令的注释
}

type People struct {
	name  string
	child *People // 结构体指针字段，类型是 *People。
}

type Address struct {
	Province    string
	City        string
	ZipCode     int
	PhoneNumber string
}

func main() {

	// 1.基本的实例化形式
	// 结构体本身是一种类型，可以像整型、字符串等类型一样，以 var 的方式声明结构体即可完成实例化。
	var p Point
	p.X = 10
	p.Y = 20

	// 2.创建指针类型的结构体

	// 可以使用 new 关键字对类型（包括结构体、整型、浮点数、字符串等）进行实例化，结构体在实例化后会形成指针类型的结构体。
	// 在 C/C++ 语言中，使用 new 实例化类型后，访问其成员变量时必须使用->操作符。
	// 在Go语言中，访问结构体指针的成员变量时可以继续使用.
	// 这是因为Go语言为了方便开发者访问结构体指针的成员变量，使用了语法糖（Syntactic sugar）技术
	// 将 p1.X 形式转换为 (*p1).X

	p1 := new(Point)
	p1.X = 5

	// 3.取结构体的地址实例化

	// 对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作，取地址格式如下：
	version := 10
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "show version"

	fmt.Println(cmd.Var) // 0xc0000ae008

	// 取地址实例化是最广泛的一种结构体实例化方式，可以使用函数封装上面的初始化过程，代码如下：
	cmd2 := newCommand("version",
		&version,
		"show version")

	fmt.Println(cmd2.Var) // 0xc0000ae008

	// 4.初始化结构体的成员变量

	// 结构体实例化后字段的默认值是字段类型的默认值，例如 ，数值为 0、字符串为 ""（空字符串）、布尔为 false、指针为 nil 等。
	relation := &People{
		name: "爷爷",
		child: &People{ // 使用取地址初始化一个 People
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}
	fmt.Println(relation.child.name)

	// 多个值列表初始化结构体, 每一个初始值的填充顺序必须与字段在结构体中的声明顺序一致
	addr := Address{
		"四川",
		"成都",
		610000,
		"0",
	}
	fmt.Println(addr) // {四川 成都 610000 0}

	// 5.初始化匿名结构体

	/*

			匿名结构体定义格式和初始化写法:

			ins := struct {
		    // 匿名结构体字段定义
		    字段1 字段类型1
		    字段2 字段类型2
		    …
			}{
		    // 字段值初始化
		    初始化字段1: 字段1的值,
		    初始化字段2: 字段2的值,
		    …
			}

	*/

	// 实例化一个匿名结构体
	msg := &struct { // 定义部分
		id   int
		data string
	}{ // 值初始化部分
		1024,
		"net error",
	}
	// 匿名结构体的类型名是结构体包含字段成员的详细描述，匿名结构体在使用时需要重新定义，造成大量重复的代码，因此开发中较少使用。
	printMsgType(msg)

	// 6.Go语言构造函数
	// Go语言的类型或结构体没有构造函数的功能，但是我们可以使用结构体初始化的过程来模拟实现构造函数。

}

// 多种方式创建和初始化结构体 ——- 模拟构造函数重载

type Cat struct {
	Color string
	Name  string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}
func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

// 带有父子关系的结构体的构造和初始化 ——- 模拟父级构造调用
type BlackCat struct {
	// 嵌入了 Cat 结构体，BlackCat 拥有 Cat 的所有成员，实例化后可以自由访问 Cat 的所有成员。
	Cat // 嵌入Cat, 类似于派生
}

// “构造基类”
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// “构造子类”
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{} // 实例化 BlackCat 结构，此时 Cat 也同时被实例化。
	cat.Color = color
	return cat
}

// 打印消息类型, 传入匿名结构体
func printMsgType(msg *struct {
	id   int
	data string
}) {
	// 使用动词%T打印msg的类型
	fmt.Printf("%T\n", msg) // *struct { id int; data string }
}

// 取地址实例化是最广泛的一种结构体实例化方式
func newCommand(name string, varref *int, comment string) *Command {

	return &Command{
		Name:    name,
		Var:     varref,
		Comment: comment,
	}
}
