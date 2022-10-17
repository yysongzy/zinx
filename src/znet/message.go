package znet

import "zinx/src/ziface"

type Message struct {
	DataLen uint32
	MsgId   uint32
	Data    []byte
}

func NewMessage(msgId uint32, data []byte) ziface.IMessage {
	m := &Message{
		DataLen: uint32(len(data)),
		MsgId:   msgId,
		Data:    data,
	}

	return m
}

func (m *Message) SetDataLen(dataLen uint32) {
	m.DataLen = dataLen
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) SetMsgId(msgId uint32) {
	m.MsgId = msgId
}

func (m *Message) GetMsgId() uint32 {
	return m.MsgId
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) GetData() []byte {
	return m.Data
}
