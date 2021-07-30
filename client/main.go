package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	coon, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("客服端连接失败，错误信息： ", err)
		return
	}
	for {
		//发送数据
		_, err = coon.Write([]byte("client -------------->"))
		if err != nil {
			fmt.Println("客户端写入数据失败，错误信息 ", err)
			return
		}

		//读取服务器发来的数据
		buf := make([]byte, 512)
		n, _ := coon.Read(buf)
		fmt.Printf("从服务器读到的数据： %s ,长度为%v 字节 \n", buf, n)

		time.Sleep(time.Second * 2)
	}
}
