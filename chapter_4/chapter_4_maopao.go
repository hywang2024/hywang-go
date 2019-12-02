package main

import "fmt"

func main() {
	arr := [...]int{12, 31, 2, 33, 31, 44, 521}
	len := len(arr)
	fmt.Println("--------没排序前--------\n", arr)
	for i := 0; i < len-1; i++ {
		for j := i; j < len-1; j++ {
			if arr[i] > arr[j] {
				t := arr[i]
				arr[i] = arr[j]
				arr[j] = t
			}
			fmt.Println(arr)
		}
	}
	fmt.Println("--------最终结果--------\n", arr)
}
