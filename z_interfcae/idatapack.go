package z_interfcae

/*
	封包，拆包，面向TCP连接中的数据流，解决TCP粘包问题
*/

type IDataPack interface {
	//获取包的头的长度
	GetHeadLen() uint32
	//封包方法
	Pack(msg IMessage) ([]byte, error)
	//拆包方法
	UnPack([]byte) (IMessage, error)
}
