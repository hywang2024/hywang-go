package main

import "fmt"

func calc(a, b int) int {
	//默认情况下会将 c 和 x 分配在栈上 ,函数结束时，栈上内存会释放
	var c int
	c = a * b
	var x int
	x = c * 10
	return x
}

// 本函数测试入口参数和返回值情况
func dummy(b int) int {
	// 声明一个变量c并赋值
	var c int
	c = b
	return c
}

// 空函数, 什么也不做
func void() {
}
func main() {
	// 声明a变量并打印
	/*var a int
	// 调用void()函数
	void()
	// 打印a变量的值和dummy()函数返回
	fmt.Println(a, dummy(0))*/

	fmt.Println(dummy2())
}

//go run -gcflags "-m -l" *.go
//go run -gcflags "-m -l" chapter_2_stack.go

/*
# command-line-arguments
.\chapter_2_stack.go:31:13: main ... argument does not escape
.\chapter_2_stack.go:31:13: a escapes to heap
.\chapter_2_stack.go:31:22: dummy(0) escapes to heap
0 0
*/

// 声明空结构体测试结构体逃逸情况
type Data struct {
}

func dummy2() *Data {
	// 实例化c为Data类型
	var c Data
	//返回函数局部变量地址
	return &c
}
