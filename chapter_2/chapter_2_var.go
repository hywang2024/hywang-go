package main

import "fmt"

func main() {
	//变量的声明方式
	//标准格式
	var str string = "这是一个字符串"
	fmt.Println(str)
	//批量格式
	var (
		a int    //默认值是 0
		b string //默认值是 空
		c byte
		d float32
		f bool
		e struct {
			x int
			y string
		}
	)
	fmt.Println(a, b, c, d, f, e)
	// 简短格式
	st := "简短定义"
	fmt.Println(st)

	//简单的赋值
	m := 100
	n := 200
	m, n = n, m
	fmt.Println(m, n)

	//数组的初始化
	var arr [3]int = [3]int{2, 4, 6}
	fmt.Println(len(arr))

}

type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}
func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
