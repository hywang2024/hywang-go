package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	valueByKey = make(map[string]string)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

func main() {
	fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("5")
	fmt.Println("---------------------------------------")

	readValue("1")

}

// 根据键读取值
func readValue(key string) string {
	// 对共享资源加锁
	valueByKeyGuard.Lock()
	// 取值
	v := valueByKey[key]
	// 对共享资源解锁
	valueByKeyGuard.Unlock()
	// 返回值
	return v
}

func readFileSize(fileName string) int64 {
	// 根据文件名打开文件, 返回文件句柄和错误
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("文件读取失败")
		return 0
	}
	//   取文件状态信息
	info, err := file.Stat()
	if err != nil {
		fmt.Println("文件信息读取失败")
		return 0
	}
	//延迟调用Close, 此时Close不会被调用
	defer file.Close()
	return info.Size()
}
