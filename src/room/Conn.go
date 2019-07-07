package room

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	pb "goServer/src/pb"
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
	return tc
}

//接收消息
func (c *TcpConn)ReadMessage() {
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

func (c *TcpConn)Close(){
	c.conn.Close()
}