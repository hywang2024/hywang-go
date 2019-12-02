package main

import (
	"fmt"
	"math"
)

//声明全局变量
var p int
var va int = 20

func main() {
	//变量的声明方式
	//标准格式
	var str string = "这是一个字符串"
	fmt.Println(str)
	//批量格式
	var (
		a int    //默认值是 0
		b string //默认值是 空
		c byte
		d float32
		f bool
		e struct {
			x int
			y string
		}
	)
	fmt.Println(a, b, c, d, f, e)
	// 简短格式
	st := "简短定义"
	fmt.Println(st)

	//简单的赋值
	m := 100
	n := 200
	m, n = n, m
	fmt.Println(m, n)

	//数组的初始化
	var arr [3]int = [3]int{2, 4, 6}
	fmt.Println(len(arr))

	i, i2 := getData()
	fmt.Println(i, i2)
	// 匿名变量
	i3, _ := getData()
	fmt.Println(i3)
	_, i4 := getData()
	fmt.Println(i4)

	//局部变量
	var x int = 2
	var y int = 3
	z := x + y

	fmt.Printf("x=%d,y=%d,z=%d \n", x, y, z)

	var va string = "局部变量和全部变量名称可以相同，但是函数会优先考虑局部变量"
	fmt.Printf("va=%v \n", va)

	i5 := sum(1, 2)

	fmt.Printf("i5 = %d", i5)

	//整型
	var va1 int = 2
	var va2 int8 = 2
	var va3 int16 = 2
	//类型不同 是无法复制 作比较
	fmt.Printf("va1 =%d,va2 =%d ,va3 = %d \n", va1, va2, va3)
	// 浮点
	var fl float32 = 16777216 // 1 << 24
	fmt.Println(fl == fl+1)   // "true"
	//可以只写 小数部分 或者 整数部分
	const c1 = .71828 // 0.71828
	const c2 = 1.     // 1
	fmt.Println(c1, c2)
	//科学计数法
	const Avogadro = 6.02214129e23 // 阿伏伽德罗常数
	const Planck = 6.62606957e-34  // 普朗克常数
	//“%f”来控制保留几位小数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
	// 复数
	var co1 complex128 = complex(1, 2) // 1+2i
	var co2 complex128 = complex(3, 4) // 3+4i
	fmt.Println(co1)
	fmt.Println(co2)
	fmt.Println(co1 * co2)       // "(-5+10i)"
	fmt.Println(real(co1 * co2)) // "-5"
	fmt.Println(imag(co1 * co2)) // "10"
}
func sum(a, b int) int {
	fmt.Printf("sum() 函数中形参 a = %d\n", a)
	fmt.Printf("sum() 函数中形参 b = %d\n", b)
	return a + b
}

func getData() (int, int) {
	return 100, 200
}
