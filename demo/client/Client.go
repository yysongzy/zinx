package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/src/znet"
)

func main() {
	fmt.Println("this is client")

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Dial error: ", err)
		return
	}

	var msgId uint32 = 0

	for {
		msgId = (msgId + 1) % 2

		dp := znet.NewDataPack()
		packMessage, err := dp.Pack(znet.NewMessage(msgId, []byte("Client Data Pack")))
		if err != nil {
			fmt.Println("Pack error: ", err)
			continue
		}

		if _, err := conn.Write(packMessage); err != nil {
			fmt.Println("Write error: ", err)
			continue
		}

		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, headData); err != nil {
			fmt.Println("ReadFull headData error: ", err)
			continue
		}

		msgHead, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("UnPack error: ", err)
			continue
		}
		if msgHead.GetDataLen() > 0 {
			msg, ok := msgHead.(*znet.Message)
			if ok {
				msg.Data = make([]byte, msg.GetDataLen())
				if _, err := io.ReadFull(conn, msg.Data); err != nil {
					fmt.Println("ReadFull data error: ", err)
					continue
				}
				fmt.Println("msgId = ", msg.GetMsgId(), ", msgData = ", string(msg.GetData()))
			} else {
				fmt.Println("Transfer failed")
			}
		}

		time.Sleep(3 * time.Second)
	}
}
