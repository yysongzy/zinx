package znet

import (
	"fmt"
	"strconv"
	"zinx/src/ziface"
)

type MsgHandler struct {
	APIs map[uint32]ziface.IRouter
}

func NewMsgHandler() ziface.IMsgHandler {
	return &MsgHandler{
		APIs: make(map[uint32]ziface.IRouter),
	}
}

func (mh *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	handler, ok := mh.APIs[request.GetMsgId()]
	if !ok {
		fmt.Println("no handler, msgId", request.GetMsgId())
		return
	}

	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (mh *MsgHandler) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := mh.APIs[msgId]; ok {
		panic("repeated msgId, msgId: " + strconv.Itoa(int(msgId)))
	}

	mh.APIs[msgId] = router
	fmt.Println("add msgId, msgId: ", msgId)
}
