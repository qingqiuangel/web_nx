package znet

import "zinx/z_interfcae"

type Request struct {
	coon z_interfcae.IConnection
	msg  z_interfcae.IMessage
}

//得到当前链接
func (r *Request) GetConnection() z_interfcae.IConnection {
	return r.coon
}

//得到请求的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
