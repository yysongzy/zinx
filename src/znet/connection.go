package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/src/ziface"
)

type Connection struct {
	Conn           *net.TCPConn
	ConnID         uint32
	isClosed       bool
	MsgHandler     ziface.IMsgHandler
	ExitBufferChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) *Connection {
	c := &Connection{
		Conn:           conn,
		ConnID:         connID,
		isClosed:       false,
		MsgHandler:     msgHandler,
		ExitBufferChan: make(chan bool, 1),
	}

	return c
}

func (c *Connection) StartReader() {
	fmt.Println("StartReader")
	defer fmt.Println("StartReader exit")
	defer c.Stop()

	for {
		dp := NewDataPack()

		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("ReadFull headData error: ", err)
			c.ExitBufferChan <- true
			continue
		}

		msg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("UnPack error: ", err)
			c.ExitBufferChan <- true
			continue
		}

		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("ReadFull data error: ", err)
				c.ExitBufferChan <- true
				continue
			}
		}
		msg.SetData(data)

		req := Request{
			conn: c,
			msg:  msg,
		}

		go func(request ziface.IRequest) {
			c.MsgHandler.DoMsgHandler(request)
		}(&req)
	}
}

func (c *Connection) Start() {
	go c.StartReader()

	for {
		select {
		case <-c.ExitBufferChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.isClosed == true {
		return
	}

	c.isClosed = true

	c.Conn.Close()

	c.ExitBufferChan <- true

	close(c.ExitBufferChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection is closed")
	}

	dp := NewDataPack()
	msg, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("Pack error: ", err)
		return errors.New("Pack error")
	}

	if _, err := c.Conn.Write(msg); err != nil {
		fmt.Println("Write error: ", err)
		return errors.New("Write error")
	}

	return nil
}
