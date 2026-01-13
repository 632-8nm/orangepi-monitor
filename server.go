package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// DataResponse 定义返回给前端的 JSON 结构
// 使用 `json:"..."` 标签可以避免你在图片里看到的那个“命名不一致”的坑
type DataResponse struct {
	Temperature string  `json:"cpu_temp"`  // 前端看到的是 cpu_temp
	Usage       float64 `json:"cpu_usage"` // 前端看到的是 cpu_usage
}

// StartServer 启动 Web 服务
func StartServer(port string) {
	// 接口 1：返回 JSON 数据
	http.HandleFunc("/api/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		res := DataResponse{
			Temperature: GetCPUTemp(),
			Usage:       GetCPUUsage(),
		}
		json.NewEncoder(w).Encode(res)
	})

	// --- 新增：托管静态网页 ---
	// 当访问根目录 http://localhost:8080/ 时，自动寻找并展示 index.html
	http.Handle("/", http.FileServer(http.Dir("./")))

	fmt.Printf("Web 服务已启动，监听端口 %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("启动失败: %v\n", err)
	}
}
