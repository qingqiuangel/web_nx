package main

import (
	"fmt"
	"zinx/z_interfcae"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(request z_interfcae.IRequest) {
	fmt.Println("当前业务PingRouter Handle")
	fmt.Printf("receive msg ID = %d ,receive data = %s \n", int(request.GetMsgID()), string(request.GetData()))
	err := request.GetConnection().SendMsg(0, []byte("服务器写给client0"))
	if err != nil {
		fmt.Println("Handle中服务器回写失败 ,err = ", err)
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (p *HelloRouter) Handle(request z_interfcae.IRequest) {
	fmt.Println("当前业务HelloRouter Handle")
	fmt.Printf("receive msg ID = %d ,receive data = %s \n", int(request.GetMsgID()), string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte(" 服务器写给client1------------------->"))
	if err != nil {
		fmt.Println("Handle中服务器回写失败 ,err = ", err)
	}
}

func main() {
	server := znet.NewServer()
	//给当前zinx 框架 添加一个路由
	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &HelloRouter{})
	server.Server()
}
