package main

import (
	"fmt"
	"strings"
)

func main() {

	//定义一个数组
	arr := []string{"go scanner", "go parser", "go compiler", "go printer", "go formater"}
	//定义处理链
	chain := []func(string) string{
		removePreFix,
		strings.TrimSpace,
		strings.ToUpper,
	}
	out := process(arr, chain)
	fmt.Println(out)
}

func removePreFix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func process(arr []string, chain []func(string) string) []string {
	for index, value := range arr {
		result := value
		for _, ch := range chain {
			result = ch(result)
		}
		arr[index] = result
	}
	return arr
}
