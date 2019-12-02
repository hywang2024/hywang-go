package main

import (
	"fmt"
)

func main() {
	//定义一个长度为10的数组
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//取第0个元素
	fmt.Printf("arr第0个元素是 %d\n", arr[0])
	//取最后一个元素
	fmt.Printf("arr最后一个元素是 %d \n", arr[len(arr)-1])
	//...”省略号，则表示数组的长度是根据初始化值的个数来计算
	arr1 := [...]int{1, 2, 3}
	// 打印索引和元素
	for i, v := range arr1 {
		fmt.Printf("第%d个元素值是 %d\n", i, v)
	}
	//比较两个数组是否相等
	// 数组长度不一样，不能直接比较
	//fmt.Println(arr==arr1)
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 2}
	fmt.Println(arr1 == arr2, arr1 == arr3)

	arr4 := [...]string{"a", "b", "c"}
	for k, v := range arr4 {
		fmt.Println(k, v)
	}
	// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
	//var array [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	//array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化数组中索引为 1 和 3 的元素
	//array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 声明并初始化数组中指定的元素
	//array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	//二维数组
	//arr5:=[...][...]int{{1,2},{1}}

}
