package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			fmt.Printf("%d %d \n", x, y)
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}
	// 手动返回, 避免执行进入标签
	return
	// 标签
breakHere:
	fmt.Println("done")

	err := firstCheckError()
	if err != nil {
		goto onExit
	}
	err = secondCheckError()
	if err != nil {
		goto onExit
	}
	fmt.Println("done")
	return
onExit:
	fmt.Println(err)
	exitProcess()
}

func exitProcess() {
}

func secondCheckError() interface{} {
	return false
}

func firstCheckError() interface{} {
	return false
}
