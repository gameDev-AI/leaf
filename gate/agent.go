package gate

import (
	"net"
)

type Agent interface {
	WriteMsg(msg interface{})
	WriteMsgWithCode(opCode, msg interface{})
	WriteMsgWithReturn(msg interface{}) [][]byte
	WriteMsgDirect([][]byte)
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
}
