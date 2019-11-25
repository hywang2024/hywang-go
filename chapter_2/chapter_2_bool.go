package main

import "fmt"

func main() {
	var a = 10
	fmt.Println(a == 5)

	// &&的优先级比||高（&& 对应逻辑乘法，|| 对应逻辑加法，乘法比加法优先级要高）
	if 1 == 2 || 2 == 2 && 2 == 2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if 'a' <= c && c <= 'z' ||
		'A' <= c && c <= 'Z' ||
		'0' <= c && c <= '9' {
		// ...ASCII字母或数字...
	}
}
