package main

import (
	"fmt"
)

// 将NewInt定义为int类型
type NewInt int

// int 起一个别名
type IntAlias = int

func main() {
	var i1 NewInt
	var i2 IntAlias

	fmt.Printf("i1的默认值是%d,类型是 %T\n", i1, i1)
	fmt.Printf("i2的默认值是%d,类型是 %T\n", i2, i2)
}
