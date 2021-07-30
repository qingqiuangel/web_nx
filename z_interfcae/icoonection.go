package z_interfcae

import "net"

type IConnection interface {
	//启动链接 让当前的链接准备开始工作
	Start()
	//停止链接 结束当前链接的工作
	Stop()
	//获取当前链接的绑定 socket coon
	GetTCPConnection() *net.TCPConn
	//获取当前链接的ID
	GetConnID() uint32
	//获取远程客服端的TCP 状态 IP Port
	RemoteAddr() net.Addr
	//发送数据
	SendMsg(uint32, []byte) error
}

//定义一个处理连接业务的方法

type HandleFunc func(*net.TCPConn, []byte, int) error
