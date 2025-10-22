package main

import (
	"fmt"
	"net/http"
)

// healthzHandler 是一个简单的 HTTP 处理函数，用于健康检查
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func main() {
	// 注册 /healthz 路由，使用单独定义的 healthzHandler 函数
	http.HandleFunc("/healthz", healthzHandler)

	fmt.Println("Room server listening on :8080")
	// ListenAndServe 应该监听 Docker 容器内的地址，nil 表示使用默认多路复用器
	err := http.ListenAndServe(":8080", nil)

	// 增加错误处理，这是良好的 Go 编程习惯
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
