package z_interfcae

/*
	消息管理抽象层
*/

type IMsgHandle interface {
	//调度/执行对应的 路由消息 处理方法
	DoMsgHandler(request IRequest)

	AddRouter(msgId uint32, router IRouter)

	CreatWorkerPool ()

	SendMsgToTaskQueue(request IRequest)
}
