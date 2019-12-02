package main

import (
	"fmt"
	"sync"
)

func main() {
	var mp map[string]string
	//赋值1
	//直接赋值是会报错，因为map需要初始化才能使用
	//mp["key1"]="value1"
	//mp["key2"]="value2"
	//赋值2
	mp2 := map[string]string{"key3": "value3", "key4": "value4"}

	//make
	mp1 := make(map[int]string)
	mp1[1] = "1"

	//遍历
	for k, v := range mp {
		fmt.Println(k, v)
	}
	// 删除
	delete(mp2, "key3")

	for k, v := range mp2 {
		fmt.Println(k, v)
	}

	//map 不能并发 读写
	var smap sync.Map
	//sync.Map 可以直接使用，不用初始化
	smap.Store("Key1", "Value")
	smap.Store("1", "2")
	smap.Store(2, 2)
	//取元素
	fmt.Println(smap.Load("Key1"))
	fmt.Println(smap.Load(1))
	//删除
	smap.Delete("1")
	//遍历
	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
