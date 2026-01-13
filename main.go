package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Orange Pi 实时监控系统后端 ===")

	// 启动服务器，监听 8080 端口
	// 这个函数在 server.go 中定义
	StartServer(":8080")
}
