package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义了框架使用的请求处理函数签名
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 是我们框架的核心引擎，它将实现 http.Handler 接口
type Engine struct {
	router map[string]HandlerFunc
}

// New 是 Engine 的构造函数
func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

// addRoute 是一个内部方法，用于把路由注册到 map 里
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern // 例如 "GET-/hello"
	engine.router[key] = handler
}

// GET 暴露给用户的 GET 请求注册方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 暴露给用户的 POST 请求注册方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动 HTTP 服务器
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 核心！实现这个方法，Engine 就能接管所有的 HTTP 请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		// 如果字典里有这个路由，就交给对应的函数处理
		handler(w, req)
	} else {
		// 没找到就返回 404
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL.Path)
	}
}
