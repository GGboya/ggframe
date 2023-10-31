package main

import (
	"ggframe/framework"
	"net/http"
)

func main() {
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: framework.NewCore(),
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
