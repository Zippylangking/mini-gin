package main

import (
	"fmt"

	"github.com/Zippylangking/mini-gin/gee"

	"net/http"
)

func main() {
	// 1. 初始化我们的框架引擎
	r := gee.New()

	// 2. 注册路由
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "欢迎来到 Mini-Gin! URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	// 3. 启动服务，监听 9999 端口
	fmt.Println("Server is running at http://localhost:9999")
	r.Run(":9999")
}
