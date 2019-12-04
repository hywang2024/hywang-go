package main

import (
	"errors"
	"fmt"
)

func main() {
	var err = errors.New("this is an error")
	fmt.Println(err)

	var e error
	// 创建一个错误实例，包含文件名和行号
	e = newParseError("main.go", 1)
	// 通过error接口查看错误描述
	fmt.Println(e.Error())
	// 根据错误接口具体的类型，获取详细错误信息
	switch detail := e.(type) {
	case *ParseError: // 这是一个解析错误
		fmt.Printf("Filename: %s Line: %d\n", detail.name, detail.line)
	default: // 其他类型的错误
		fmt.Println("other error")
	}
}

//声明一个解析错误
type ParseError struct {
	name string
	line int
}

//实现error接口
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s : %d \n", e.name, e.line)
}

func newParseError(fileName string, line int) error {
	return &ParseError{fileName, line}
}
