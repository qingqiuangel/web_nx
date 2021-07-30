package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/utils"
	"zinx/z_interfcae"
)

type Connection struct {
	//指定当前的server
	Server z_interfcae.IServer
	//当前连接的socket Tcp套接字
	Coon *net.TCPConn

	//连接的ID
	CoonID uint32

	//当前的连接状态
	isClosed bool

	//是否退出的 channel
	ExitChan chan bool

	//消息管道  Reader--->Writer
	MsgChan chan []byte

	//该链接的处理方法Router
	MsgHandler z_interfcae.IMsgHandle
}

func NewConnection(server z_interfcae.IServer,tcpCoon *net.TCPConn, CoonId uint32, handle z_interfcae.IMsgHandle) *Connection {
	c := &Connection{
		Server:		server,
		Coon:       tcpCoon,
		CoonID:     CoonId,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		MsgChan:    make(chan []byte),
		MsgHandler: handle,
	}
	c.Server.GetConnectionManager().AddConnection(c)
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine ....start] ")
	defer fmt.Println("[Reader is exit ] coon ID =", c.CoonID, " ,remote addr is ", c.RemoteAddr().String())
	defer c.Stop()
	for {
		dp := &DataPack{}
		//读取客服端二进制流 head msg 的前8个字节
		headData := make([]byte, dp.GetHeadLen())
		_, err := c.GetTCPConnection().Read(headData)
		if err != nil {
			fmt.Println("读取客服端二进制流 head msg 的前8个字节失败", err)
			return
		}
		//做拆包处理
		msg, err := dp.UnPack(headData)
		//根据dataLen 读取真正的data 数据
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err = io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("拆包读取data失败", err)
				break
			}
		}
		msg.SetData(data)
		//得到Request数据
		request := &Request{
			coon: c,
			msg:  msg,
		}
		//go c.MsgHandler.DoMsgHandler(request)
		 if utils.GlobalObject.WorkerPoolSize >0 {
			 c.MsgHandler.SendMsgToTaskQueue(request)
		 }else {
		 	panic("工作池未开启！！！")
		 }
	}
}
func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine ....start]")
	defer fmt.Println("[Writer is exit ]coon ID =", c.CoonID, " ,remote addr is ", c.RemoteAddr().String())
	for {
		select {
		case data := <-c.MsgChan:
			if _, err := c.Coon.Write(data); err != nil {
				fmt.Println("send message failed err", err)
				return
			}
		case <-c.ExitChan:
			return
		}
	}
}

//启动链接 让当前的链接准备开始工作
func (c *Connection) Start() {
	//启动读数据业务
	go c.StartReader()
	//启动写数据业务
	go c.StartWriter()

}

//停止链接 结束当前链接的工作
func (c *Connection) Stop() {
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	c.ExitChan <- true
	c.Coon.Close()
	c.Server.GetConnectionManager().DeleteConnection(c)
	close(c.MsgChan)
	close(c.ExitChan)
}

//获取当前链接的绑定 socket coon
func (c *Connection) GetTCPConnection() *net.TCPConn {
	TcpConn := c.Coon
	return TcpConn
}

//获取当前链接的ID
func (c *Connection) GetConnID() uint32 {
	id := c.CoonID
	return id
}

//获取远程客服端的TCP 状态 IP Port
func (c *Connection) RemoteAddr() net.Addr {
	addr := c.Coon.RemoteAddr()
	return addr
}

//发送数据
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("无法向客服端发送数据，链接已关闭")
	}
	dp := NewDataPack()
	BinaryMsg, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("向客户端打包数据 ，pack出错", err)
		return err
	}
	c.MsgChan <- BinaryMsg
	return nil
}
