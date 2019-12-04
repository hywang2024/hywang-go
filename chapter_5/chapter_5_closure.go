package main

import "fmt"

func main() {
	i := 1
	i2 := incr(i)
	fmt.Printf("%d \n", i2())
	fmt.Printf("%d \n", i2())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", i2)

	// 创建一个累加器, 初始值为1
	accumulator2 := incr(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator2)

	generator := playerGen("high noon")
	name, hp := generator()
	fmt.Println(name, hp)
}

func incr(i int) func() int {
	// 返回一个闭包
	return func() int {
		i++
		return i
	}
}

func playerGen(name string) func() (string, int) {
	hp := 100
	return func() (s string, i int) {
		return name, hp
	}
}
