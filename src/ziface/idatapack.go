package ziface

type IDataPack interface {
	GetHeadLen() uint32
	Pack([]byte) (IMessage, error)
	UnPack(IMessage) ([]byte, error)
}
