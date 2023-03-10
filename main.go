package main

import (
	"net/http"
	"os"

	"gohade/core-practise-demo/framework"
)

func main() {
	server := &http.Server{
		// 请求监听地址
		Addr: ":8080",
		// 自定义的请求核心处理函数
		Handler: framework.NewCore(),
	}

	if err := server.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}
