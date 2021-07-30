package znet

import (
	"zinx/z_interfcae"
)

type BaseRouter struct {
}

/*
	这里所有的BaseRouter 方法都为空，因为有的Router不希望有PreHandle或者其他方法，
	BaseRouter实现了 Irouter接口的全部方法，
*/

//处理coon业务之前的钩子方法Hook
func (b *BaseRouter) PreHandle(request z_interfcae.IRequest) {

}

//处理coon业务的主方法hook
func (b *BaseRouter) Handle(request z_interfcae.IRequest) {
}

//处理coon业务之后的方法hook
func (b *BaseRouter) PostHandle(request z_interfcae.IRequest) {
}
