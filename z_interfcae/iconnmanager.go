package z_interfcae
 
/*
	链接管理模块 抽象层 
 */

type IConnManager interface {
		//添加链接
		AddConnection(IConnection)
		//删除链接
		DeleteConnection(IConnection)
		//根据conn Id 获取链接
		GetConnectionById(connId uint32) (IConnection,error)
		//获取链接个数
		GetConnectionLength() int
		//清除所有链接
		ClearAllConnection()
}