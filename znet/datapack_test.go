package znet

import (
	"fmt"
	"net"
	"testing"
)

func TestDataPacK(t *testing.T) {
	//模拟服务器，创建socket连接
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("监听端口失败，err ", err)
		return
	}
	//创建一个go ，承载处理客服端业务
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("listener.Accept()，err ", err)
				return
			}
			d := NewDataPack()
			for {
				//拆包的过程，第一次读取head中的DataLen
				/*		headData := make([]byte, d.GetHeadLen())
						reader := bufio.NewReader(conn)
						_,err := reader.Read(headData)
						if err != nil {
							fmt.Println("io 读取头信息失败，err ", err)
							break
						}*/
				headData := make([]byte, d.GetHeadLen())
				_, err := conn.Read(headData)
				if err != nil {
					fmt.Println("io 读取头信息失败，err ", err)
					break
				}
				iMessage, err := d.UnPack(headData)
				if err != nil {
					fmt.Println("d.UnPack(headData)失败，err ", err)
					return
				}
				fmt.Println("---->")
				if iMessage.GetMsgLen() > 0 {
					//第二次从coon ，根据长度再次读取data
					/*		msg := iMessage.(*Message)
							msg.Data = make([]byte, msg.GetMsgLen())*/
					//根据dataLen 的长度，继续从io流中读取
					msgLen := iMessage.GetMsgLen()
					id := iMessage.GetMsgID()
					data := make([]byte, int(msgLen))
					_, err = conn.Read(data)
					if err != nil {
						fmt.Println("根据dataLen 的长度，继续从io流中失败，err ", err)
						return
					}
					fmt.Printf("完整的消息读取完毕，msg ID 为%d,datalen = %d ,数据内容为%s \n", id, msgLen, string(data))
				}
			}
		}
	}()
	/*
		模拟客户端
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("客户端链接失败，err ", err)
		return
	}
	dp := NewDataPack()
	//封装第一个包
	msg1 := &Message{
		DataLen: 5,
		Id:      1,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}

	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("sendData1失败，err ", err)
		return
	}
	//封装第二个包
	msg2 := &Message{
		DataLen: 3,
		Id:      2,
		Data:    []byte{'g', 'y', 'l'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("sendData2失败，err ", err)
		return
	}
	//讲两个包粘在一起
	sendData1 = append(sendData1, sendData2...)
	conn.Write(sendData1)

}
