package main

import "fmt"

func main() {
	arr := [...]int{12, 14, 22, 25, 26, 27, 30, 34, 35, 37, 66, 99, 100}
	find := BinaryFind(&arr, 0, len(arr)-1, 300)
	fmt.Println(find)
}

//二分查找函数 //假设有序数组的顺序是从小到大（很关键，决定左右方向）
func BinaryFind(arr *[13]int, leftIndex int, rightIndex int, findVal int) bool {
	//判断 leftIndex是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return false
	}
	//先找到 中间的下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		//要查找的数，范围应该在 leftIndex 到 middle+1
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		//要查找的数，范围应该在 middle+1 到 rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标为：%v \n", middle)
		return true
	}
	return false
}
