package room

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"goServer/src/logger"
	pb "goServer/src/pb"
	"net"
	"sync"
)

var (
	count uint64
)

type TcpConn struct {
	sync.Mutex
	Addr            string
	MaxConnNum      int
	PendingWriteNum int
	serverid        uint64
	remoteAddr      net.Addr
	conn            net.Conn
	listener        net.Listener
}

func CreateConn(c net.Conn) *TcpConn {
	tc := &TcpConn{}
	tc.serverid = count
	tc.conn = c
	count += 1
	tc.remoteAddr = tc.conn.RemoteAddr()
	return tc
}

func (c *TcpConn) InitTcpConn() {
	var err error
	c.listener, err = net.Listen("tcp", c.Addr)
	if err != nil {
		logger.Fatal_MSG("%v", err)
	}
	if c.MaxConnNum <= 0 {
		c.MaxConnNum = 100
	}
	if c.PendingWriteNum <= 0 {
		c.PendingWriteNum = 100
	}
}

func (c *TcpConn) run() {
	//var tempDelay time.Duration
	r := CreateRoom()
	for {
		conn, err := c.listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {

			}
		}
		pl := CreateConn(conn)
		r.AddConn(pl)
		if err != nil {
			panic(err)
		}
		fmt.Println("new connect", conn.RemoteAddr())
		go pl.ReadMessage()
	}
}

//接收消息
func (c *TcpConn) ReadMessage() {
	buf := make([]byte, 409600)
	for {
		len := 0
		//读消息
		cnt, err := c.conn.Read(buf)
		if err != nil {
			continue
		}
		stReceive := &pb.UserInfo{}
		fmt.Println(cnt, buf[0:2])
		for len < cnt {
			res1 := int(buf[len+1])
			res2 := int(buf[len] << 8)
			res := res1 + res2
			pData := buf[len+2 : len+res+2]
			err = proto.Unmarshal(pData, stReceive)
			if err != nil {
				panic(err)
			}
			fmt.Println("receive", c.conn.RemoteAddr(), stReceive)
			len += 12
		}
	}
}

func (c *TcpConn) Close() {
	c.conn.Close()
}

func (c *TcpConn) Destroy() {
	c.Lock()
	defer c.Unlock()
}
