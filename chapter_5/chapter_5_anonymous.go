package main

import (
	"fmt"
)

func main() {
	// 匿名函数
	func(data int) {
		fmt.Println("匿名函数传入的值是", data)
	}(100)

	i := func(data int) int {
		return data
	}
	fmt.Println("匿名函数赋值给i", i(10))

	arr := []int{1, 2, 2, 3, 23}
	//匿名函数 作为回调
	visit(arr, func(i int) {
		fmt.Println(i)
	})

}
func visit(arr []int, f func(i int)) {
	for _, value := range arr {
		f(value)
	}
}
