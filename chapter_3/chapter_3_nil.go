package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//默认值
	var m map[int]string
	var ptr *int
	var c chan int
	var sl []int
	var f func()
	var i interface{}
	fmt.Printf("%#v\n", m)
	fmt.Printf("%#v\n", ptr)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", sl)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%#v\n", i)
	//不同类型的nil 占用的内存不一样
	var p *struct{}
	fmt.Println(unsafe.Sizeof(p)) // 8
	var s []int
	fmt.Println(unsafe.Sizeof(s)) // 24
	var mp map[int]bool
	fmt.Println(unsafe.Sizeof(mp)) // 8
	var ch chan string
	fmt.Println(unsafe.Sizeof(ch)) // 8
	var fc func()
	fmt.Println(unsafe.Sizeof(fc)) // 8
	var in interface{}
	fmt.Println(unsafe.Sizeof(in)) // 16
}
