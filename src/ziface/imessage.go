package ziface

type IMessage interface {
	SetDataLen(uint32)
	GetDataLen() uint32

	SetMsgId(uint32)
	GetMsgId() uint32

	SetData([]byte)
	GetData() []byte
}
