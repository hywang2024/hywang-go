package main

import "fmt"

func main() {
	var arr [20]int
	for i := 0; i < 20; i++ {
		arr[i] = i + 1
	}
	//区间
	fmt.Println(arr[10:13])
	fmt.Println(arr[10:])
	fmt.Println(arr[:13])
	fmt.Println(arr[:])

	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片
	var numListEmpty = []int{}
	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)

	//make Type 是指切片的元素类型，size 指的是为这个类型分配多少个元素，cap 为预分配的元素数量，
	// 这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题。
	arr1 := make([]int, 10, 30)
	arr2 := make([]int, 10)
	fmt.Println(cap(arr2))
	fmt.Println(arr1, arr2)
	fmt.Println(len(arr1), len(arr2))
	//append 追加元素

	arr1 = append(arr1, 2, 3, 3, 33, 3)
	fmt.Println(arr1)

	for i := 0; i < 30; i++ {
		arr1 = append(arr1, i)
		fmt.Printf("len: %d,cap: %d ,pointer %p \n", len(arr1), cap(arr1), arr1)
	}

	//添加元素
	arr3 := make([]int, 5)
	fmt.Println(arr3)
	arr3 = append(arr1, arr3...)
	fmt.Println(arr3)

}
