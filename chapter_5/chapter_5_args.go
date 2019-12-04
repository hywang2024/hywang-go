package main

import (
	"bytes"
	"fmt"
)

func main() {
	arg(1, 2, 3)

	args("1", 2)

	fmt.Println("a", "b")
	fmt.Println("a", "b", "a", "b", "a", "b")

	print("a", "b", "a", "b", "a", "b")
}

// 相同类型参数
func arg(arg ...int) {
	fmt.Println(arg)
}

// 不同类型参数
func args(args ...interface{}) {
	fmt.Println(args)
}

func joinStr(slist ...string) string {
	// 定义一个字节缓冲, 快速地连接字符串
	var b bytes.Buffer
	for _, value := range slist {
		// 将遍历出的字符串连续写入字节数组
		b.WriteString(value)
	}
	return b.String()
}

func printValueType(value ...interface{}) string {
	var typeStr string
	var b bytes.Buffer

	for _, val := range value {
		str := fmt.Sprintf("%v", val)
		switch val.(type) {
		case bool:
			typeStr = "bool"
		case string: // 当s为字符串类型时
			typeStr = "string"
		case int: // 当s为整型类型时
			typeStr = "int"
		}
		b.WriteString(str)
		b.WriteString("的类型是")
		b.WriteString("的类型是")
	}
	return typeStr
}
func rawPrint(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}
func print(args ...interface{}) {
	rawPrint(args)
}
