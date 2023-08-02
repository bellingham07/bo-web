package boServer

import (
	"bo-web/boContext"
	"bo-web/routes"
	"fmt"
	"net/http"
)

type Server interface {
	// Route 路由，命中路由会执行handle func
	// method get put delete
	Route(method string, pattern string, handleFunc func(ctx *boContext.Context))
	// Start 启动服务器
	Start(addr string) error
}

// sdkHttpServer 基于net/http实现的http server
type sdkHttpServer struct {
	// Name server的名字，给个标记，日志输出用
	Name    string
	handler *routes.HandlerBasedOnMap
}

func (s *sdkHttpServer) Route(method, pattern string, handleFunc func(ctx *boContext.Context)) {
	// 第一种方法
	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	//	ctx := boContext.NewContext(writer, request)
	//	handleFunc(ctx)
	//})

	// 第二种方法
	// 只注册一次，所以需要放到Start里面去

	// 让handler自己注册
	key := s.handler.Key(method, pattern)
	s.handler.Handlers[key] = handleFunc
}

func (s *sdkHttpServer) Start(addr string) error {
	http.Handle("/", s.handler)
	return http.ListenAndServe(addr, nil)
}

// 返回一个实际类型所需要的指针的时候，是需要一个指针的
func NewHttpServer(name string) Server {
	return &sdkHttpServer{Name: name}
}

// 在没有 context 抽象的情况下，是长这样的
func SignUp(ctx *boContext.Context) {
	req := &signUpReq{}

	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
	}

	// 返回虚拟id
	resp := &commonResponse{Data: 123}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入响应失败：%v", err)
	}

	// 返回一个虚拟的 user id 表示注册成功了
	//fmt.Fprintf(w, "%d", 123)
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
