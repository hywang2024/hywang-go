package main

import (
	"flag"
	"fmt"
)

//定义命令行参数
var mode = flag.String("mode", "", "process mode")

func main() {
	flag.Parse()
	fmt.Println(*mode)

	str := new(string)
	*str = "Go语言教程"
	fmt.Println(*str)
}

//go run main.go --mode=fast
