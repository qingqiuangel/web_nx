package znet

type Message struct {
	DataLen uint32
	Id      uint32
	Data    []byte
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		DataLen: uint32(len(data)),
		Id:      id,
		Data:    data,
	}
}

//获取消息ID
func (m *Message) GetMsgID() uint32 {
	return m.Id
}

//获取消息的长度
func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

//获取消息的内容
func (m *Message) GetData() []byte {
	return m.Data
}

//设置消息的ID
func (m *Message) SetMsgID(id uint32) {
	m.Id = id
}

//设置消息的长度
func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len

}

//设置消息的内容
func (m *Message) SetData(data []byte) {
	m.Data = data
}
