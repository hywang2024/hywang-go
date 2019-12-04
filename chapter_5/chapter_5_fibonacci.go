package main

import (
	"fmt"
	"time"
)

//斐波那契数列
func main() {
	now := time.Now()
	for i := 0; i < 40; i++ {
		i2 := fibonacci2(i)
		fmt.Println(i2)
	}
	sub := time.Now().Sub(now)
	fmt.Println("耗时", sub)
}

//1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, …
func fibonacci(n int) int {
	if n < 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func Factorial(n int) int {
	if n > 0 {
		return n * Factorial(n-1)
	} else {
		return 0
	}
}

const LIM = 41

var fibs [LIM]uint64

func fibonacci2(n int) (res uint64) {
	// 记忆化：检查数组中是否已知斐波那契（n）
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	fibs[n] = res
	return
}
