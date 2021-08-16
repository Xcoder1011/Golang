package main

import (
	"fmt"
)

// Go语言map的多键索引
//多个数值条件可以同时查询

type Model struct {
	Name    string
	Age     int
	Address string
}

type QueryKey struct {
	Name string
	Age  int
}

func main() {

	list := []*Model{
		{Name: "张三", Age: 23, Address: "上海"},
		{Name: "李四", Age: 25},
		{Name: "王五"},
	}

	// 传统查询
	findData(list, "张三", 23) // 输出：  &{张三 23 上海}

	//利用map的多键索引查询（组合键查询）
	buildIndex(list)    //构建基于查询的组合键（name、age)
	queryData("张三", 23) //依据name、age进行查询（多条件）

}

var mapper = make(map[QueryKey]*Model)

func buildIndex(list []*Model) {
	for _, model := range list {
		key := QueryKey{
			Name: model.Name,
			Age:  model.Age,
		}

		mapper[key] = model
	}

}

// 多键索引
func queryData(name string, age int) {

	key := QueryKey{Name: name, Age: age}

	result, ok := mapper[key]
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("没有找到对应的数据")
	}
}

// 传统查询
func findData(list []*Model, name string, age int) {
	for _, model := range list {
		if model.Name == name && model.Age == age {
			fmt.Println(model)
			return
		}
	}
	fmt.Println("没有找到对应的数据")
}
