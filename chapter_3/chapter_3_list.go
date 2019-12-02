package main

import (
	"container/list"
	"fmt"
)

func main() {
	//list 定义
	li := list.New()
	var li1 list.List

	//添加元素
	li.PushBack(1)
	li1.PushFront(2)
	fmt.Printf("%d", li.Front())
	fmt.Printf("%d", li1.Front())

	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	// 移除
	l.Remove(element)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
