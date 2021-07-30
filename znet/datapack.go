package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"zinx/utils"
	"zinx/z_interfcae"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

//获取包的头的长度
func (d *DataPack) GetHeadLen() uint32 {
	//DataLen uint32(4字节) +ID uint32 (4字节)
	return 8
}

//封包方法
//DataLen | msgID | data
func (d *DataPack) Pack(msg z_interfcae.IMessage) ([]byte, error) {
	//创建一个存放bytes字节的缓冲区
	buf := new(bytes.Buffer)
	//把DataLen 写入缓冲区
	err := binary.Write(buf, binary.LittleEndian, msg.GetMsgLen())
	if err != nil {
		fmt.Println("把DataLen 写入缓冲区 失败， err信息", err)
		return nil, err
	}
	//把MsgID   写入缓冲区
	err = binary.Write(buf, binary.LittleEndian, msg.GetMsgID())
	if err != nil {
		fmt.Println("把MsgID 写入缓冲区 失败， err信息", err)
		return nil, err
	}
	//将Data 数据写入缓冲区
	err = binary.Write(buf, binary.LittleEndian, msg.GetData())
	if err != nil {
		fmt.Println("将Data写入缓冲区 失败， err信息", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

//拆包方法
func (d *DataPack) UnPack(binaryData []byte) (z_interfcae.IMessage, error) {
	//创建一个 二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)
	//只解压head信息， 得到DataLen 和MsgId
	msg := &Message{}
	binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen)
	binary.Read(dataBuff, binary.LittleEndian, &msg.Id)
	//判断DataLen 是否超过了全局变量设置的最大包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("单个数据包 超过了全局变量设置的最大包长度")
	}
	return msg, nil
}
