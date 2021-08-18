package main

import (
	"fmt"
)

func main() {

	arr := [...]int{21, 32, 12, 33, 34, 34, 87, 24}

	var count = len(arr)

	fmt.Println("--------没排序前--------\n", arr)

	for i := 0; i < count-1; i++ {
		fmt.Println("--------第", i+1, "次冒泡--------")
		for j := 0; j < count-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
		fmt.Println(arr)

	}
	fmt.Println("--------最终结果--------\n", arr)
}
