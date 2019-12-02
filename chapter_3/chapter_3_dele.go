package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c", "d", "e", "f", "g"}
	//要删除的index
	index := 4
	fmt.Println("删除前的数组", arr)
	arr1 := arr[:index]
	arr2 := arr[index+1:]
	fmt.Println("切除的两个数组", arr1, arr2)
	arr3 := append(arr1, arr2...)
	fmt.Println("拼接起来", arr3)

}
