package main

import (
	"fmt"
)

// 其它语言中的容器:

// C++ 语言的容器通过标准库提供，如 vector 对应数组，list 对应双链表，map 对应映射等。
// C# 语言通过 .NET 框架提供，如 List 对应数组，LinkedList 对应双链表，Dictionary 对应映射。
// Lua 语言的 table 实现了数组和映射的功能，Lua 语言默认没有双链表支持。

func main() {

	// Go语言数组的声明
	// var 数组变量名 [元素数量]Type

	var arr [3]int               // 定义三个整数的数组, 每个元素都初始化为0
	fmt.Println(arr[0])          // 打印第一个元素
	fmt.Println(arr[len(arr)-1]) // 打印最后一个元素

	// 打印索引和元素
	for i, v := range arr {
		fmt.Printf("i: %d  v:%d\n", i, v)
	}
	/*
		i: 0  v:0
		i: 1  v:0
		i: 2  v:0
	*/

	// 仅打印元素
	for _, v := range arr {
		fmt.Printf("v:%d\n", v)
	}
	/*
		v:0
		v:0
		v:0
	*/

	// 用一组值来初始化数组
	var vecotor1 [3]int = [3]int{1, 2, 3}
	var vecotor2 [3]int = [3]int{1, 2}
	fmt.Println(vecotor1[2]) // "3"
	fmt.Println(vecotor2[2]) // "0"

	// “...”省略号，则表示数组的长度是根据初始化值的个数来计算
	vecotor3 := [...]int{1, 2, 3}
	fmt.Printf("vecotor3: %T\n", vecotor3) // vecotor3: [3]int

	q := [3]int{1, 2, 3}
	// q = [4]int{1, 2, 3, 4} // 编译错误：无法将 [4]int 赋给 [3]int
	fmt.Println(q) // [1 2 3]

	// 数组类型相同（包括数组的长度，数组中元素的类型）
	// 才可以直接通过较运算符（==和!=）来判断两个数组是否相等，
	// 只有当两个数组的所有元素都是相等的时候数组才是相等的
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"

	// d := [3]int{1, 2}
	// fmt.Println(a == d) // 编译错误：无法比较 [2]int == [3]int

	arrayTest()

	sliceTest()

	sliceCopy()
}

/// 多维数组

func arrayTest() {

	// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
	var array1 [4][2]int

	// 设置每个元素的整型值
	array1[0][0] = 1
	array1[0][1] = 2
	array1[1][0] = 2
	array1[3][1] = 5

	// 使用数组字面量来声明并初始化一个二维整型数组
	array1 = [4][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}

	// 声明并初始化数组中索引为 1 和 3 的元素
	array1 = [4][2]int{0: {1, 2}, 3: {4, 5}}

	fmt.Print(array1) //  [[1 2] [0 0] [0 0] [4 5]]

}

// 切片 slice
// 切片（slice）是对数组的一个连续片段的引用，所以切片是一个引用类型
// 因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型）
// 这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集
// 终止索引标识的项不包括在切片内。

func sliceTest() {

	// Go语言中切片的内部结构包含地址、大小和容量
	// 切片默认指向一段连续内存区域，可以是数组，也可以是切片本身。

	var a = [3]int{1, 2, 3}

	// 格式如下： slice [开始位置 : 结束位置]
	// 语法说明如下：
	// slice：表示目标切片对象；
	// 开始位置：对应目标切片对象的索引；
	// 结束位置：对应目标切片的结束索引。

	var slice1 = a[1:2] // 不包含结束位置2
	var slice2 = a[0:2] // 不包含结束位置2

	fmt.Printf("\nslice1 = %v", slice1)   // slice1 = [2]
	fmt.Printf("\nslice2 = %v\n", slice2) // slice2 = [1 2]

	// 从数组或切片生成新的切片拥有如下特性：
	// 取出的元素数量为：结束位置 - 开始位置；
	// 取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)-1] 获取；
	// 当缺省开始位置时，表示从连续区域开头到结束位置；
	// 当缺省结束位置时，表示从开始位置到整个连续区域末尾；
	// 两者同时缺省时，与切片本身等效；
	// 两者同时为 0 时，等效于空切片，一般用于切片复位。

	var b [10]int
	for i := 0; i < 10; i++ {
		b[i] = i + 1
	}

	fmt.Println(b)      // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(b[:5])  // [1 2 3 4 5]
	fmt.Println(b[5:8]) // [6 7 8]
	fmt.Println(b[8:])  // [9 10]

	// 表示原有的切片
	fmt.Println(b[:]) // [1 2 3 4 5 6 7 8 9 10]
	// 重置切片，清空拥有的元素
	fmt.Println(b[0:0]) // []

	// 切片类型声明格式如下： var name []Type
	// 其中 name 表示切片的变量名，Type 表示切片对应的元素类型。

	// 声明字符串切片
	var strSlice []string

	// 声明整型切片
	var numSlice []int

	// 声明一个空切片
	var numEmptySlice = []int{}

	// 输出3个切片
	fmt.Println(strSlice, numSlice, numEmptySlice) // [] [] []

	// 输出3个切片大小
	fmt.Println(len(strSlice), len(numSlice), len(numEmptySlice)) // 0 0 0

	// 切片判定空的结果
	// 切片是动态结构，只能与 nil 判定相等，不能互相判定相等。
	fmt.Println(strSlice == nil) // true
	fmt.Println(numSlice == nil) // true
	// numEmptySlice 已经被分配到了内存，但没有元素，因此和 nil 比较时是 false。
	fmt.Println(numEmptySlice == nil) // false

	// 使用 make() 函数动态地构造切片
	// 格式如下： make( []Type, size, cap )
	// 其中 Type 是指切片的元素类型，size 指的是为这个类型分配多少个元素，
	// cap 为预分配的元素数量，这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题。
	// 切片长度 len 并不等于切片的容量 cap。

	slice3 := make([]int, 10)
	slice4 := make([]int, 10, 11)
	fmt.Println(slice3, slice4)           // [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(slice3), len(slice4)) // 10 10

	numSlice = append(numSlice, 1)
	numSlice = append(numSlice, 2, 3, 4)
	fmt.Println(numSlice)                          // [1 2 3 4]
	numSlice = append(numSlice, []int{5, 6, 7}...) // 追加一个切片, 切片需要解包
	fmt.Println(numSlice)                          // [1 2 3 4 5 6 7]

	// 扩容
	var limitSlice []int
	for i := 0; i < 10; i++ {
		limitSlice = append(limitSlice, i+1)
		fmt.Printf("len: %d cap: %d, pointer: %p\n", len(limitSlice), cap(limitSlice), limitSlice)
	}

	/*
		 1. 如果空间不足以容纳足够多的元素，切片就会进行“扩容”，此时新切片的长度会发生改变
		 2. 切片在扩容时，容量的扩展规律是按容量的 2 倍数进行扩充，例如 1、2、4、8、16……，
		 3. 使用函数 cap() 查看切片的容量情况。
		 4. 内存地址也会发生变化

		长度	   容量	   内存地址
		len: 1 cap: 1, pointer: 0xc000014338
		len: 2 cap: 2, pointer: 0xc000014340
		len: 3 cap: 4, pointer: 0xc00001e0a0
		len: 4 cap: 4, pointer: 0xc00001e0a0
		len: 5 cap: 8, pointer: 0xc0000182c0
		len: 6 cap: 8, pointer: 0xc0000182c0
		len: 7 cap: 8, pointer: 0xc0000182c0
		len: 8 cap: 8, pointer: 0xc0000182c0
		len: 9 cap: 16, pointer: 0xc00007e000
		len: 10 cap: 16, pointer: 0xc00007e000

	*/

	fmt.Println(limitSlice) // [1 2 3 4 5 6 7 8 9 10]
	// 还可以在切片的开头添加元素：
	limitSlice = append([]int{0}, limitSlice...)          // 在开头添加1个元素0
	fmt.Println(limitSlice)                               // [1 2 3 4 5 6 7 8 9 10]
	limitSlice = append([]int{-3, -2, -1}, limitSlice...) // 在开头添加1个切片
	fmt.Println(limitSlice)                               // [-3 -2 -1 0 1 2 3 4 5 6 7 8 9 10]
	limitSlice = append([]int{99}, limitSlice[3:]...)
	fmt.Println(limitSlice)                                                      //  [99 0 1 2 3 4 5 6 7 8 9 10]
	limitSlice = append(limitSlice[:3], append([]int{98}, limitSlice[3:]...)...) // 链式操作
	fmt.Println(limitSlice)                                                      // [99 0 1 98 2 3 4 5 6 7 8 9 10]

	// 从开头位置删除
	// 方法① a = a[N:]  // 删除开头N个元素
	var slice = []int{1, 2, 3, 4}
	// slice = slice[1:]   // 删除开头1个元素
	// fmt.Println(slice) // [2 3 4]
	// slice = slice[N:]   // 删除开头N个元素

	// 方法② 使用append 删除开头N个元素 （不会导致内存空间结构的变化）
	// slice = append(slice[:0], slice[1:]...)
	// slice = append(slice[:0], slice[N:]...)
	// fmt.Println(slice) // [2 3 4]

	// 方法③ 使用 copy()  删除开头N个元素
	// slice = slice[:copy(slice, slice[1:])]
	// slice = slice[:copy(slice, slice[N:])]
	// fmt.Println(slice) // [2 3 4]

	// 从中间位置删除
	slice = append(slice[:1], slice[1+1:]...)
	fmt.Println(slice) // [1 3 4]

	/*
		slice = append(slice[:i], slice[i+1:]...) // 删除中间1个元素(第i个元素)
		slice = append(slice[:i], slice[i+N:]...) // 删除中间N个元素
		slice = slice[:i+copy(slice[i:], slice[i+1:])] // 删除中间1个元素
		slice = slice[:i+copy(slice[i:], slice[i+N:])] // 删除中间N个元素
	*/

	// 从尾部删除
	slice = slice[:len(slice)-1] // 删除尾部1个元素
	// slice = slice[:len(slice) - N] 	// 删除尾部N个元素

}

/// copy()：切片复制（切片拷贝）

func sliceCopy() {

	// copy() 函数的使用格式如下： copy( destSlice, srcSlice []T) int

	// srcSlice 为数据来源切片，destSlice 为复制的目标（也就是将 srcSlice 复制到 destSlice）
	// 来源和目标的类型必须一致
	// copy() 函数的返回值表示实际发生复制的元素个数。

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1)        // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice1, slice2) // [1 2 3 4 5] [1 2 3]

	copy(slice1, slice2)        // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1, slice2) // [1 2 3 4 5] [1 2 3]

	// 设置元素数量为1000
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据， 切片不会因为等号操作进行元素的复制。
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0]) // 999
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1]) // 0  999
	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i]) // 4 5 2 3 4
	}

}
