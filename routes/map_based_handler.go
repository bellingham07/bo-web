package routes

import (
	"bo-web/boContext"
	"net/http"
)

type Routable interface {
	Route(method string, pattern string, handleFunc func(ctx *boContext.Context))
}

type Handler interface {
	// 处理请求
	ServeHTTP(c *boContext.Context)
	// 处理路由
	Routable
}

type HandlerBasedOnMap struct {
	// key = method +url
	Handlers map[string]func(ctx *boContext.Context)
}

func (h *HandlerBasedOnMap) Route(method string, pattern string, handleFunc func(ctx *boContext.Context)) {
	key := h.Key(method, pattern)
	h.Handlers[key] = handleFunc
}

func (h *HandlerBasedOnMap) ServeHTTP(c *boContext.Context) {
	key := h.Key(c.R.Method, c.R.URL.Path)
	// 找到对应路径
	if handler, ok := h.Handlers[key]; ok {
		handler(c)
	} else {
		// 没找到
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found"))
	}
}

func (h *HandlerBasedOnMap) Key(method, pattern string) string {
	return method + "#" + pattern
}

// 一种常用的go设计模式
// 用于保证HandlerBasedOnMap肯定实现了这个接口
var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		Handlers: make(map[string]func(c *boContext.Context)),
	}
}
