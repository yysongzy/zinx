package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/src/utils"
	"zinx/src/ziface"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	return 8
}

func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	dataBuf := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuf.Bytes(), nil
}

func (dp *DataPack) UnPack(binaryData []byte) (ziface.IMessage, error) {
	dataBuf := bytes.NewReader(binaryData)

	msg := &Message{}

	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.MsgId); err != nil {
		return nil, err
	}

	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large package size")
	}

	return msg, nil
}
