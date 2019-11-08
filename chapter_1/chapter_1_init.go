package main

import "fmt"

func main() {
	fmt.Println(str)
}

//非局部变量，map类型，且已初始化
var m = map[int]string{1: "一", 2: "二", 3: "三", 4: "四"}

//非局部变量，string类型，未被初始化
var str string

func init() {
	// 格式化打印
	fmt.Printf("Map: %v\n", m)
}
