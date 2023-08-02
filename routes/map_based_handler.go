package routes

import (
	"bo-web/boContext"
	"net/http"
)

type HandlerBasedOnMap struct {
	// key = method +url
	Handlers map[string]func(ctx *boContext.Context)
}

func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.Key()
	// 找到对应路径
	if handler, ok := h.Handlers[key]; ok {
		handler(boContext.NewContext(writer, request))
	} else {
		// 没找到
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Not Found"))
	}
}

func (h *HandlerBasedOnMap) Key(method, pattern string) string {
	return method + "#" + pattern
}
