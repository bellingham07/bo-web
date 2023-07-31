package server

import "net/http"

type Server interface {
	// Route 路由，命中路由会执行handle func
	Route(pattern string, handleFunc http.HandlerFunc)
	// Start 启动服务器
	Start(addr string) error
}

// sdkHttpServer 基于net/http实现的http server
type sdkHttpServer struct {
	// Name server的名字，给个标记，日志输出用
	Name string
}
