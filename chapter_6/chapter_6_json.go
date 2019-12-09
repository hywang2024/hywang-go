package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := genJson()
	fmt.Println(string(jsonData))

	//只需要json部分数据
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}
	// 反序列化到screenAndTouch
	json.Unmarshal(jsonData, &screenAndTouch)
	fmt.Println(screenAndTouch)
	fmt.Printf("%v \n", screenAndTouch)
}

// 定义一个手机屏幕
type Screen struct {
	//手机 屏幕尺寸
	Size float32
	//分辨率
	Resx, ResY int
}

// 定义电池
type Battery struct {
	// 容量
	Capacity int
}

//生成json
func genJson() []byte {
	raw := &struct {
		Screen
		Battery
		HasTouchID bool // 序列化时添加的字段：是否有指纹识别
	}{
		Screen: Screen{
			Size: 7,
			Resx: 1080,
			ResY: 1090,
		},
		Battery:    Battery{Capacity: 100},
		HasTouchID: true,
	}
	bytes, _ := json.Marshal(raw)
	return bytes
}
