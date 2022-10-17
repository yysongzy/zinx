package main

import (
	"fmt"
	"zinx/src/ziface"
	"zinx/src/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	fmt.Println("receive from client: msgId = ", request.GetMsgId(), ", data = ", string(request.GetData()))
	err := request.GetConnection().SendMsg(0, []byte("ping"))
	if err != nil {
		fmt.Println("Handle error: ", err)
		return
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (hr *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloRouter Handle")
	fmt.Println("receive from client: msgId = ", request.GetMsgId(), ", data = ", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("hello"))
	if err != nil {
		fmt.Println("Handle error: ", err)
		return
	}
}

func main() {
	s := znet.NewServer("Eric's Server")
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Serve()
}
