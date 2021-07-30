package z_interfcae

/*
	路由抽象接口，路由里的数据都是IRequest
*/

type IRouter interface {
	//处理coon业务之前的钩子方法Hook
	PreHandle(request IRequest)
	//处理coon业务的主方法hook
	Handle(request IRequest)
	//处理coon业务之后的方法hook
	PostHandle(request IRequest)
}
