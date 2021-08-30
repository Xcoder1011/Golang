package main

import (
	"fmt"
)

/*
	递归函数

构成递归需要具备以下条件：
一个问题可以被拆分成多个子问题；
拆分前的原问题与拆分后的子问题除了数据规模不同，但处理问题的思路是一样的；
不能无限制的调用本身，子问题需要有退出递归状态的条件。

*/
func main() {

	/// 1) 斐波那契数列

	// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, …

	for i := 1; i <= 10; i++ {
		result := fibonacci(i)

		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}

	/*

		打印：

		fibonacci(1) is: 1
		fibonacci(2) is: 1
		fibonacci(3) is: 2
		fibonacci(4) is: 3
		fibonacci(5) is: 5
		fibonacci(6) is: 8
		fibonacci(7) is: 13
		fibonacci(8) is: 21
		fibonacci(9) is: 34
		fibonacci(10) is: 55

	*/

	/// 2) 数字阶乘

	var i int = 10
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i))) //打印： 10 的阶乘是 3628800
	i = 3
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i))) //打印： 3 的阶乘是 6

}

/// 1) 斐波那契数列

// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, …

func fibonacci(n int) (result int) {
	if n <= 2 {
		result = 1
	} else {
		result = fibonacci(n-1) + fibonacci(n-2)
	}
	return result
}

/// 2) 数字阶乘

/// 例如，n!=1×2×3×…×n，阶乘亦可以递归方式定义：0!=1，n!=(n-1)!×n。

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
