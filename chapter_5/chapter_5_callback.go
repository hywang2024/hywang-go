package main

import (
	"flag"
	"fmt"
)

//定义命令行参数， 从命令行输入 --skill 可以将=后的字符串传入 skillParam 指针变量。
var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("fire")
		},
		"run": func() {
			fmt.Println("run")
		},
		"fly": func() {
			fmt.Println("fly")
		},
	}
	//skillParam 是一个 *string 类型的指针变量，使用 *skillParam 获取到命令行传过来的值，
	// 并在 map 中查找对应命令行参数指定的字符串的函数。
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("not found")
	}

	// go chapter_5_callback.go --skill fly
}
