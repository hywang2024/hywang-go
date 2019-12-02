package main

import "fmt"

func main() {

	for j := 0; j < 5; j++ {
	JLoop:
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop
			}
			fmt.Println(i)
		}

	}
	fmt.Println("------------------")

	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}
	fmt.Println(step)
}
