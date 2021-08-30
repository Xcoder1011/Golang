package main

import (
	"fmt"
	"time"
)

/*
	递归函数

构成递归需要具备以下条件：
一个问题可以被拆分成多个子问题；
拆分前的原问题与拆分后的子问题除了数据规模不同，但处理问题的思路是一样的；
不能无限制的调用本身，子问题需要有退出递归状态的条件。


递归函数的缺点就是比较消耗内存，而且效率比较低，通过在内存中缓存并重复利用缓存从而避免重复执行相同计算的方式

*/

func main() {

	/// 1) 斐波那契数列(普通版)

	// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, …

	var result uint64 = 0
	start := time.Now()
	for i := 1; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("数列第 %d 位: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub((start))
	fmt.Printf("斐波那契数列(普通版)的执行时间为: %s\n", delta)

	/*

		打印：

		数列第 1 位: 1
		数列第 2 位: 1
		数列第 3 位: 2
		....

		数列第 16 位: 987
		数列第 17 位: 1597
		数列第 18 位: 2584
		数列第 19 位: 4181
		数列第 20 位: 6765
		斐波那契数列(普通版)的执行时间为: 248.166µs

	*/

	/// 2) 斐波那契数列(优化版)

	start = time.Now()
	for i := 1; i < LIM; i++ {
		result = fibonacciCache(i)
		fmt.Printf("数列第 %d 位: %d\n", i, result)
	}
	end = time.Now()
	delta = end.Sub((start))
	fmt.Printf("斐波那契数列(优化版)的执行时间为: %s\n", delta)

	/*

		打印：

		数列第 1 位: 1
		数列第 2 位: 1
		数列第 3 位: 2
		....

		数列第 16 位: 987
		数列第 17 位: 1597
		数列第 18 位: 2584
		数列第 19 位: 4181
		数列第 20 位: 6765
		斐波那契数列(普通版)的执行时间为: 96.858µs

		/// 总结： 通过内存缓存显著提升了性能

	*/

	/// 3) 数字阶乘

	var i int = 10
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i))) //打印： 10 的阶乘是 3628800
	i = 3
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i))) //打印： 3 的阶乘是 6

}

/// 1) 斐波那契数列(普通版)

func fibonacci(n int) (result uint64) {
	if n <= 2 {
		result = 1
	} else {
		result = fibonacci(n-1) + fibonacci(n-2)
	}
	return result
}

/// 2) 斐波那契数列(优化版)

const LIM = 21

var cache [LIM]uint64

func fibonacciCache(n int) (result uint64) {

	// 检查数组中是否已知斐波那契（n）
	var num = cache[n]
	if num != 0 {
		result = num
		return
	}

	if n <= 2 {
		result = 1
	} else {
		result = fibonacci(n-1) + fibonacci(n-2)
	}
	// 内存缓存
	cache[n] = result
	return result
}

/// 3) 数字阶乘

/// 例如，n!=1×2×3×…×n，阶乘亦可以递归方式定义：0!=1，n!=(n-1)!×n。

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
