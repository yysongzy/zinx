package znet

import (
	"bytes"
	"testing"
)

func TestDataPack_PackAndUnPack(t *testing.T) {
	msg := &Message{}
	msg.SetDataLen(5)
	msg.SetMsgId(1)
	msg.SetData([]byte("hello"))

	dp := NewDataPack()

	binaryData, err := dp.Pack(msg)
	if err != nil {
		t.Error("Pack failed")
	}

	unPackMsg, err := dp.UnPack(binaryData)
	if err != nil {
		t.Error("UnPack failed")
	}

	if unPackMsg.GetDataLen() != 5 {
		t.Error("GetDataLen failed")
	}

	if unPackMsg.GetMsgId() != 1 {
		t.Error("GetMsgId failed")
	}

	data := binaryData[8:]
	if bytes.Compare(data, []byte("hello")) != 0 {
		t.Error("GetData failed")
	}
}
