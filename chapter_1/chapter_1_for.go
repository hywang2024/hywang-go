package main

import "fmt"

func main() {
	//for循环，去掉了多余的{}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	/*for{
		fmt.Println("这是个无线死循环")
	}*/

	// if
	if true {
		fmt.Println("这里是true")
	} else {
		fmt.Println("这里是false")
	}
	if false {
		fmt.Println("这里是false")
	} else {
		fmt.Println("这里是true")
	}
}
