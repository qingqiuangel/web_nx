package z_interfcae

type IServer interface {
	Start()
	Stop()
	Server()
	//路由功能，给当前的服务注册一个路由方法，供客服端使用
	AddRouter(msgId uint32, router IRouter)
	//在server模块中，获取整个Connection 管理器
	GetConnectionManager() IConnManager
}
