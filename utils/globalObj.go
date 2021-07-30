package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx/z_interfcae"
)

var GlobalObject *GlobalObj

type GlobalObj struct {
	/*
		Server
	*/
	TcpServer z_interfcae.IServer //当前Zinx的全局Server对象
	Host      string              //服务器主机监听的IP
	TcpPort   int                 //端口
	Name      string              //当前服务器的名称

	/*
		Zinx
	*/
	Version        string //Zinx版本号
	MaxCoon        int    //当前服务器允许的最大连接数
	MaxPackageSize uint32 //Zinx框架数据包的最大值(buf 的大小)
	//工作池的worker 数量
	WorkerPoolSize  uint32
	//每个消息队列允许创建的最大任务数
	MaxWorkerTaskSize uint32
}

func (g *GlobalObj) LoadConfigFile() {
	data, err := ioutil.ReadFile("Version_Test/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

//提供Init方法，初始化当前GlobalObj 数据
func init() {
	//默认的值
	GlobalObject = &GlobalObj{
		Host:           "0.0.0.0",
		TcpPort:         8080,
		Name:           "ZinxServer",
		Version:        "V1.4",
		MaxCoon:        2,
		MaxPackageSize: 4096,
		WorkerPoolSize: 10,
		MaxWorkerTaskSize:1024,
	}
	GlobalObject.LoadConfigFile()
}
