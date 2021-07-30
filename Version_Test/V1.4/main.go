package main

import (
	"fmt"
	"zinx/z_interfcae"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) PreHandle(request z_interfcae.IRequest) {
	fmt.Println("当前业务PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before PING"))
	if err != nil {
		fmt.Println("PreHandle 回写错误，错误信息：", err)
	}

}
func (p *PingRouter) Handle(request z_interfcae.IRequest) {
	fmt.Println("当前业务Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before PING"))
	if err != nil {
		fmt.Println("Handle 回写错误，错误信息：", err)
	}
}
func (p *PingRouter) PostHandle(request z_interfcae.IRequest) {
	fmt.Println("当前业务PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before PING"))
	if err != nil {
		fmt.Println("PostHandle 回写错误，错误信息：", err)
	}
}

func main() {
	server := znet.NewServer()
	//给当前zinx 框架 添加一个路由
	server.AddRouter(&PingRouter{})
	server.Server()
}
