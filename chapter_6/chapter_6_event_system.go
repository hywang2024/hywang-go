package main

import "fmt"

//生成一个map 用来保存时间名称 和 时间函数 数组
var eventMap = make(map[string][]func(interface{}))

//向map中注册事件
func RegisterEvent(eventName string, eventfunc func(interface{})) {
	//先从map中取出事件数据
	funArr := eventMap[eventName]
	//在数据中 增加事件函数
	funArr = append(funArr, eventfunc)
	//将数组 放入map中
	eventMap[eventName] = funArr
}

//调用函数   interface{}可以传入 任意类型
func CallEvent(eventName string, param interface{}) {
	//获取到事件 数组
	funArr := eventMap[eventName]
	//循环事件数组，执行事件函数
	for _, fu := range funArr {
		fu(param)
	}
}

func main() {
	ac := new(actor)
	RegisterEvent("onEvent", ac.OnEvent)
	RegisterEvent("globalEvent", GlobalEvent)
	CallEvent("globalEvent", 100)
}

type actor struct {
}

func (ac actor) OnEvent(param interface{}) {
	fmt.Println("actor onevent", param)
}
func GlobalEvent(param interface{}) {
	fmt.Println("global event param ", param)
}
