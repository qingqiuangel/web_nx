package z_interfcae

type IRequest interface {
	//得到当前链接
	GetConnection() IConnection

	//得到请求的数据
	GetData() []byte

	GetMsgID() uint32
}
