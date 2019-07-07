package connection

import (
	"net"
)

type Connection struct {
	*net.TCPConn
}
