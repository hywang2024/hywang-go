package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	//图片大小
	const size = 300
	// 创建灰度图
	gray := image.NewGray(image.Rect(0, 0, size, size))

	//遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			//填充为白色
			gray.SetGray(x, y, color.Gray{255})
		}
	}
	//生成x轴坐标
	for x := 0; x < size; x++ {
		//让sin值在 0-2PI 之间
		f := float64(x) * 2 * math.Pi / size
		//sin 的幅度为一半的像素 向下便宜一半像素并翻转
		i := size/2 - math.Sin(f)*size/2
		// 用黑色 设置sin轨迹
		gray.SetGray(x, int(i), color.Gray{0})

	}
	//创建文件
	file, e := os.Create("sin.png")
	if e != nil {
		log.Fatal(e)
	}

	//将数据写入文件
	png.Encode(file, gray)

	// 关闭文件
	file.Close()
}
