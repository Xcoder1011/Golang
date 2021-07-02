package main

import (
	"fmt"
	"reflect"
)

/*

 在 Go 1.9 版本之前定义内建类型的代码是这样写的：
type byte uint8
type rune int32


而在 Go 1.9 版本之后变为：
type byte = uint8
type rune = int32

*/

// 类型别名 与 类型定义

// 类型别名: 将int取一个别名叫IntAlias
type IntAlias = int

// 类型定义: 将NewInt定义为int类型
type NewInt int

func main() {

	var a NewInt

	fmt.Printf("a type: %T\n", a) // main.NewInt

	var b IntAlias

	fmt.Printf("b type: %T\n", b) // int

	// a 的类型是 main.NewInt，表示 main 包下定义的 NewInt 类型，
	// b 类型是 int，IntAlias 类型只会在代码中存在，编译完成时，不会有 IntAlias 类型。

	testBrand()
}

// 定义商标结构
type Brand struct {
}

// 为商标结构添加Show()方法
func (t Brand) Show() {

}

// 为Brand定义一个别名MyBrand
type MyBrand = Brand

// 定义车辆结构
type Vehicle struct {
	MyBrand
	Brand
}

func testBrand() {

	// 声明变量car为车辆类型
	var car Vehicle

	// 指定调用MyBrand的Show
	car.MyBrand.Show()

	// 取car的类型反射对象
	ta := reflect.TypeOf(car) // 使用反射取变量 car 的反射类型对象，以查看其成员类型。

	// 遍历car的所有成员
	for i := 0; i < ta.NumField(); i++ { // 遍历 car 的结构体成员
		// car的成员信息
		f := ta.Field(i)
		// 打印成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.
			Name())
	}

	// FieldName: MyBrand, FieldType: Brand
	// FieldName: Brand, FieldType: Brand
}
