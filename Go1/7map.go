package main

import (
	"fmt"
)

func main() {

	// range关键字：循环迭代切片

	slice := []int{1, 2, 3, 4, 5}

	for i, v := range slice {
		fmt.Printf("i: %d  v:%d\n", i, v)
	}
	/*

		i: 0  v:1
		i: 1  v:2
		i: 2  v:3
		i: 3  v:4
		i: 4  v:5

	*/

	for index, value := range slice {

		fmt.Printf("value: %d  value地址: %X  每个元素的地址: %X\n", value, &value, &slice[index])
	}

	/*
		迭代返回的变量value是一个在迭代过程中根据切片依次赋值的新变量，所以 value 的地址总是相同的

		value: 1  value地址: C0000AE020  每个元素的地址: C0000AC060
		value: 2  value地址: C0000AE020  每个元素的地址: C0000AC068
		value: 3  value地址: C0000AE020  每个元素的地址: C0000AC070
		value: 4  value地址: C0000AE020  每个元素的地址: C0000AC078
		value: 5  value地址: C0000AE020  每个元素的地址: C0000AC080
	*/

	for index := 0; index < len(slice); index++ {
		value := slice[index]
		fmt.Printf("value: %d  value地址: %X  \n", value, &value)
	}

	multiSliceTest()

	mapTest()

}

// 多维切片

func multiSliceTest() {

	var slice [][]int // 声明一个二维切片
	slice = [][]int{{10}, {20, 30}}
	fmt.Println(slice) // [[10] [20 30]]

	slice1 := [][]int{{10}, {20, 30}} // 声明一个二维整型切片并赋值
	slice1[0] = append(slice1[0], 40) // 为第一个切片追加值为 40 的元素
	fmt.Println(slice1)               //  [[10 40] [20 30]]
}

// Go语言映射
// map: 关联数组或字典
func mapTest() {

	// 声明：var mapname map[keytype]valuetype
	// mapname 为 map 的变量名。
	// keytype 为键类型。
	// valuetype 是键对应的值类型。

	var mapList map[string]int
	mapList = map[string]int{"one": 1, "two": 2}
	mapList["three"] = 3
	fmt.Println(mapList) // map[one:1 three:3 two:2]

	map2 := make(map[string]float32) // 等价于 map2 := map[string]float{}
	map2["key1"] = 2.3
	map2["key2"] = 0.6
	fmt.Println(map2) // map[key1:2.3 key2:0.6]

	// map 容量
	// map 可以根据新增的 key-value 动态的伸缩，因此它不存在固定长度或者最大限制，
	// 可以选择标明 map 的初始容量 capacity，
	// 格式如下： make(map[keytype]valuetype, cap)

	map5 := make(map[string]string, 10)

	// 当 map 增长到容量上限的时候，如果再增加新的 key-value，map 的大小会自动加 1，
	// 出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。

	// 用切片作为 map 的值
	map3 := make(map[string][]int)
	map4 := make(map[int]*[]int)

	fmt.Println(map3, map4, map5)

	for k, v := range mapList {
		fmt.Printf("k: %v  v:%d\n", k, v)
		/*
			k: three  v:3
			k: one  v:1
			k: two  v:2
		*/
	}

	for k := range mapList {
		fmt.Printf("k: %v  v:%d\n", k, mapList[k])
	}

}
