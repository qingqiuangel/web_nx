package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/z_interfcae"
)

type Server struct {
	//服务器的名称
	Name string
	//服务器绑定的IP版本
	IpVersion string
	//服务器监听的IP地址
	Ip string
	//端口号
	Port int
	//当前Server的消息管理模块
	MsgHandler z_interfcae.IMsgHandle
	//当前Server的Connection管理模块
	ConnMgr z_interfcae.IConnManager
}

func (s *Server) Start() {
	fmt.Printf("[zinx] Server Name :%s ,listen IP %s ,Port: %d \n", utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	s.MsgHandler.CreatWorkerPool()
	go func() {
		//获取一个TCP Addr
		addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
		if err != nil {
			fmt.Println("获取一个TCP Addr 失败 err :", err)
			return
		}
		listener, err := net.ListenTCP(s.IpVersion, addr)
		var TcpConnId uint32
		TcpConnId = 0
		for {
			coon, err := listener.AcceptTCP()
			//判断当前链接数是否超出  设定的最大链接数    目前还没开辟新的链接，假如现在有10个链接，最大链接数是11.此时 新的链接已经开辟了，但是还没加到连接管理模块
			if s.ConnMgr.GetConnectionLength()+1 > utils.GlobalObject.MaxCoon {
				fmt.Println("---------------------》 超过了限定的最大连接数！《------------------")
				coon.Close()
				continue
			}
			if err != nil {
				fmt.Println("listener.AcceptTCP()失败 err ", err)
				return
			}
			c := NewConnection(s,coon, TcpConnId, s.MsgHandler)
			TcpConnId++
			go c.Start()
		}
	}()

}

func (s *Server) Stop() {
	//如果停止服务，需要释放全部的资源和链接
	s.ConnMgr.ClearAllConnection()

}

func (s *Server) Server() {
	s.Start()

	//阻塞，等待以后填充操作
	select {}
}

//初始化Server
func (s *Server) AddRouter(msgId uint32, router z_interfcae.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
	fmt.Printf("添加msgID = %v router 成功 \n", msgId)
}
func NewServer() z_interfcae.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IpVersion:  "tcp4",
		Ip:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:	NewConnManager(),
	}
	return s
}
func (s *Server) GetConnectionManager() z_interfcae.IConnManager  {
	return s.ConnMgr
}