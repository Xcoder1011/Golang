package main

// 二分查找算法

// 二分查找的核心思想是将 n 个元素分成大致相等的两部分，取中间值 a[n/2] 与 x 做比较，
// 如果 x=a[n/2]，则找到 x，算法中止，如果 x<a[n/2]，则只要在数组 a 的左半部分继续搜索 x，
// 如果 x>a[n/2]，则只要在数组 a 的右半部搜索 x。

import (
	"fmt"
)

//二分查找函数 //假设有序数组的顺序是从小到大（很关键，决定左右方向）
func BinaryFind(arr *[]int, leftIndex int, rightIndex int, findVal int) {
	//判断 leftIndex是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	//先找到 中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		//要查找的数，范围应该在 leftIndex 到 middle+1
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		//要查找的数，范围应该在 middle+1 到 rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标为：%v \n", middle)
	}
}

func main() {
	//定义一个数组
	arr := []int{1, 2, 5, 7, 15, 25, 30, 36, 39, 51, 67, 78, 80, 82, 85, 91, 92, 97}
	BinaryFind(&arr, 0, len(arr)-1, 30)
	fmt.Println("main arr=", arr)
}
