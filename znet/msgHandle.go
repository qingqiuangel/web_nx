package znet

import (
	"errors"
	"fmt"
	"zinx/utils"
	"zinx/z_interfcae"
)

type MsgHandle struct {
	//存放每个msg ID 的对应处理路由方法
	Apis map[uint32]z_interfcae.IRouter
	//任务消息队列
	TaskWorkerQueue [] chan z_interfcae.IRequest
	//工作池的worker 数量
	WorkerPoolSize  uint32
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]z_interfcae.IRouter),
		WorkerPoolSize:utils.GlobalObject.WorkerPoolSize,
		TaskWorkerQueue:make([]chan z_interfcae.IRequest,utils.GlobalObject.WorkerPoolSize),
	}
}

func (m *MsgHandle) DoMsgHandler(request z_interfcae.IRequest) {
	//从Request 中找到MsgId
	handler, ok := m.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("执行此msg ID 不存在，请添加！")
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (m *MsgHandle) AddRouter(msgId uint32, router z_interfcae.IRouter) {
	//判断map中是否存在该msg id
	if _, ok := m.Apis[msgId]; ok {
		errors.New("已经存在了此msgID ,无法再次添加！！！")
	}
	m.Apis[msgId] = router
}

func (m *MsgHandle) CreatWorkerPool (){
	for i:= 0 ; i< int(m.WorkerPoolSize) ;i++ {
		//第 i 个worker 就用 第 i 个消息队列channel
		m.TaskWorkerQueue[i] =make(chan z_interfcae.IRequest,utils.GlobalObject.MaxWorkerTaskSize)
		go m.StartOneWorkerProcess(i,m.TaskWorkerQueue[i])
	}
}

func(m *MsgHandle)StartOneWorkerProcess (i int , TaskWorkerQueue chan z_interfcae.IRequest)  {
	fmt.Println("workID = " ,i, "已经启动就绪，等待消息载入！")
	for {
		select {
		case request := <-TaskWorkerQueue :
			m.DoMsgHandler(request)
		}
	}
}
func (m *MsgHandle)  SendMsgToTaskQueue (request z_interfcae.IRequest){
	//将消息平均分配给不同的worker， 通过coonID/总数量 取余
	workerId := request.GetConnection().GetConnID() % m.WorkerPoolSize
	fmt.Println("当前处理的工作----------> workerID =",workerId)
	m.TaskWorkerQueue[workerId] <- request
}
