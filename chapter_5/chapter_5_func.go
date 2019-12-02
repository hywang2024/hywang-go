package main

import "fmt"

func main() {
	in := Data{
		complax:   []int{1, 2, 3},
		innerData: InnerData{22},
		ptr:       &InnerData{1},
	}

	// 输入结构的成员情况
	fmt.Printf("in value: %+v\n", in)
	// 输入结构的指针地址
	fmt.Printf("in ptr: %p\n", &in)
	// 传入结构体，返回同类型的结构体
	out := PassByValye(in)
	// 输出结构的成员情况
	fmt.Printf("out value: %+v\n", out)
	// 输出结构的指针地址
	fmt.Printf("out ptr: %p\n", &out)
}

type InnerData struct {
	a int
}
type Data struct {
	complax   []int     //切片类型
	innerData InnerData //实例
	ptr       *InnerData
}

func PassByValye(data Data) Data {
	//输出参数成员   //如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
	fmt.Printf("data value %+v \n", data)
	// 打印参数指针
	fmt.Printf("data pointer %p \n", &data)
	return data
}
