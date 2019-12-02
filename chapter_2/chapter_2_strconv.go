package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 100
	// 整型转换成string
	itoa := strconv.Itoa(num)
	fmt.Printf("%T \n", itoa)
	str := string(num)
	fmt.Printf("%T \n", str)

	//字符串转整型
	str1 := "100"
	i, _ := strconv.Atoi(str1)
	fmt.Printf("%d 的类型是 %T \n", i, i)
	str2 := "a1"
	i2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Printf("%v str2转换失败", i2)
	} else {
		fmt.Printf("%d 转换成功 类型是 %T \n", i2, i2)
	}
	// 布尔类型 只能接受 1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	bo := "t"
	b, err := strconv.ParseBool(bo)
	if err != nil {
		fmt.Printf("bo转换失败")
	} else {
		fmt.Printf("bo转换成功 %T \n", b)
	}

	n := "-11" // 可以接受负数
	//base 指定进制，取值范围是 2 到 36。如果 base 为 0，则会从字符串前置判断，“0x”是 16 进制，“0”是 8 进制，否则是 10 进制。
	//bitSize 指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64。
	//返回的 err 是 *NumErr 类型的，如果语法有误，err.Error = ErrSyntax，如果结果超出类型范围 err.Error = ErrRange。
	i3, err := strconv.ParseInt(n, 10, 0)
	if err != nil {
		fmt.Printf("n转换成功")
	} else {
		fmt.Printf("n转换成功 %T %d \n", i3, i3)
	}

	s := "-11" //不接受 负数
	nu, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(nu)
	}
	// 小数
	st := "3.1415926"
	num1, err := strconv.ParseFloat(st, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num1)
	}

	//格式化 format
	num2 := true
	formatBool := strconv.FormatBool(num2)
	fmt.Printf("type:%T,value:%v\n ", formatBool, formatBool)

	// append
	// 声明一个slice
	b10 := []byte("int (base 10):")

	// 将转换为10进制的string，追加到slice中
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))
	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))
}
