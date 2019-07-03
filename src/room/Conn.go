package room

import (
	"net"
)
var(
	count uint64
)

type TcpConn struct{
	connid uint64
	remoteAddr net.Addr
	conn net.Conn
}


func CreateConn(c net.Conn) *TcpConn{
	tc :=&TcpConn{}
	tc.connid =count
	tc.conn = c
	count +=1
	tc.remoteAddr = tc.conn.RemoteAddr()
}
