package main

import (
	"net/http"
)

func main() {
	// 启动一个简单的 http服务
	//文件服务器将当前目录作为根目录（/目录）的处理器，访问根目录，就会进入当前目录
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)

}
