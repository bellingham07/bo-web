package main

import (
	"bo-web/boServer"
	"fmt"
	"net/http"
)

func main() {
	server := boServer.NewHttpServer("test-server")
	server.Route("/", home)
	server.Route("/user", user)
	server.Route("/user/create", createUser)
	server.Route("/user/signup", boServer.SignUp)
	server.Route("/order", order)
	server.Start(":8080")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是订单")
}
