package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Go语言通过通道可以实现多个 goroutine 之间内存共享

	//创建一个string 通道
	channel := make(chan string)
	// 创建producer()函数的并发goroutine
	go producter("str1", channel)
	go producter("str2", channel)
	consumer(channel)
}

//生产者  注意 chan<- 的写法
func producter(header string, channel chan<- string) {
	//无线循环 生成数据
	for {
		//生成随机数，放入到channel 通道中
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		time.Sleep(time.Second)
	}

}

//消费者 注意 <-chan 的写法
func consumer(channel <-chan string) {
	//无线循环  取数据
	for {
		// 从通道中取出数据, 此处会阻塞直到信道中返回数据
		message := <-channel
		// 打印数据
		fmt.Println(message)
	}
}
