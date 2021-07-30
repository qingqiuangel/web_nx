package main

import (
	"fmt"
	"net"
	"time"
	"zinx/znet"
)

func main() {
	coon, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("客服端连接失败，错误信息： ", err)
		return
	}
	for {
		//客户端发送消息
		dp := znet.NewDataPack()
		binaryClientMsg, err := dp.Pack(znet.NewMessage(0, []byte("client0 message ....")))
		if err != nil {
			fmt.Println("client0 pack msg failed ,err= ", err)
			return
		}
		_, err = coon.Write(binaryClientMsg)
		if err != nil {
			fmt.Println("client0 coon.Write failed ,err= ", err)
			return
		}

		//服务器应该收到一个sever 中Handle 写入的消息， msgId =1 ,内容为 Ping ping ping 的消息
		//先读取流中 msg head 部分的id ,dataLen , 在二次读取数据
		headData := make([]byte, dp.GetHeadLen())
		_, err = coon.Read(headData)
		if err != nil {
			fmt.Println("client0 coon.Read failed ,err= ", err)
			return
		}
		//拆包
		headMessage, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("client0 dp.UnPack failed ,err= ", err)
			return
		}
		var data []byte
		if headMessage.GetMsgLen() > 0 {
			data = make([]byte, headMessage.GetMsgLen())
			coon.Read(data)
		}
		fmt.Println("receive server data = ", string(data))
		time.Sleep(time.Second * 2)
	}
}
