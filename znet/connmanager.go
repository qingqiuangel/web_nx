package znet

import (
	"errors"
	"fmt"
	"sync"
	"zinx/z_interfcae"
)

/*
	链接管理模块实现层
*/

type ConnManager struct {
	Connections map[uint32] z_interfcae.IConnection
	ConnLock        sync.RWMutex
}

func  NewConnManager () *ConnManager {
	return &ConnManager{
		Connections: make(map[uint32]z_interfcae.IConnection),
	}
}

func (c *ConnManager) AddConnection(connection z_interfcae.IConnection) {
	//加写锁
	c.ConnLock.Lock()
	defer c.ConnLock.Unlock()

	c.Connections[connection.GetConnID()] =connection
	fmt.Println("connection add to ConnManager successfully! connId = " ,connection.GetConnID())
}

func (c *ConnManager) DeleteConnection(connection z_interfcae.IConnection) {
	//加写锁
	c.ConnLock.Lock()
	defer c.ConnLock.Unlock()

	delete(c.Connections,connection.GetConnID())
	fmt.Println("connection delete to ConnManager successfully! connId = " ,connection.GetConnID())
}

func (c *ConnManager) GetConnectionById(connId uint32) (z_interfcae.IConnection, error) {
	//加读锁
	c.ConnLock.RLock()
	defer c.ConnLock.RUnlock()

	if conn,ok := c.Connections[connId] ; ok {
		return conn,nil
	}else {
		return nil, errors.New("GetConnectionById failed")
	}
}

func (c *ConnManager) GetConnectionLength() int {
	return  len(c.Connections)
}

func ( c *ConnManager) ClearAllConnection() {
	//加写锁
	c.ConnLock.Lock()
	defer c.ConnLock.Unlock()

	for connId, connection := range c.Connections {
		//关闭连接 一系列操作
		connection.Stop()
		delete(c.Connections,connId)
	}
}
