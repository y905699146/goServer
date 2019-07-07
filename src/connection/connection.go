package connection

import (
	"net"
)

type Connection interface{
	ReadMsg()([]byte,error)
	WriterMsg(arg ...[]byte) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
}
