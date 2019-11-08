package main

import "fmt"

func main() {
	//定义字符串
	var str = "this is string"
	println(str)
	var varInt8 int8 = 100
	var varByte byte = 'c'
	var varUint8 uint8 = 'a'

	println(varInt8)
	println(varByte)
	println(varUint8)

	//常量
	const (
		a = 1
		b = iota
		c
		d
		f
		g = 111
		h
		i
	)
	fmt.Println(a, b, c, d, f, g, h, i)

	type color byte
	const (
		black color = iota
		red
		blue
		green
	)

	println("Enum color: ", black, red, blue, green)

	const (
		_, _ = iota, iota * 10
		first, firstPrize
		snd, sndPrize
		third, thirdPrize
		forth, forthPrize
	)
	println("Enum rank prize:", first, firstPrize)
	println("Enum rank prize:", snd, sndPrize)
	println("Enum rank prize:", third, thirdPrize)
	println("Enum rank prize:", forth, forthPrize)

	// 切片
	ints := make([]int, 0, 5)

	println(ints, "Slice")

	//map
	strings := make(map[string]string)
	strings["s1"] = "s1"
	strings["s2"] = "s2"
	println(strings["s1"])

	const x = 123
	const y = 1.23
	fmt.Println(y)
}
