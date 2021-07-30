package z_interfcae

/*
	将请求的消息封装到一个message中
*/

type IMessage interface {
	//获取消息ID
	GetMsgID() uint32
	//获取消息的长度
	GetMsgLen() uint32
	//获取消息的内容
	GetData() []byte

	//设置消息的ID
	SetMsgID(uint32)
	//设置消息的长度
	SetDataLen(uint32)
	//设置消息的内容
	SetData([]byte)
}
