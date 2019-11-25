package main

import "fmt"

func main() {
	const pi = 3.14
	const (
		p = 3
		f = 4
	)

	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)

	const (
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagBlue
	)
	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)
	//二进制
	fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue)

	fmt.Printf("%s %d", GPU, GPU)
}

// 声明芯片类型
type ChipType int

const (
	None ChipType = iota
	CPU           // 中央处理器
	GPU           // 图形处理器
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}
