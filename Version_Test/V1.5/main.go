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
	fmt.Println("当前业务Handle")
	fmt.Printf("receive msg ID = %d ,receive data = %s \n", int(request.GetMsgID()), string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("ping。。。ping....Ping。。。"))
	if err != nil {
		fmt.Println("Handle中服务器回写失败 ,err = ", err)
	}
}

func main() {
	server := znet.NewServer()
	//给当前zinx 框架 添加一个路由
	server.AddRouter(&PingRouter{})
	server.Server()
}
