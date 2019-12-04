package main

import (
	"fmt"
	"time"
)

func main() {
	test()
}

func test() {
	now := time.Now()
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum++
	}
	//since := time.Since(now)
	sub := time.Now().Sub(now)
	fmt.Println(sum, "test func() 耗时", sub)
}
