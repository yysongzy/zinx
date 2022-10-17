package ziface

import "net"

type IConnection interface {
	StartReader()
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() uint32
	GetRemoteAddr() net.Addr
	SendMsg(msgId uint32, data []byte) error
}
