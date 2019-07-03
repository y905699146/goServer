package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "goServer/src/pb"
	"goServer/src/room"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:6600")
	if err != nil {
		panic(err)
	}
	r  :=room.CreateRoom()
	for {
		conn, err := listener.Accept()
		pl :=room.CreateConn(conn)
		r.AddConn(pl)
		if err != nil {
			panic(err)
		}
		fmt.Println("new connect", conn.RemoteAddr())
		go readMessage(conn)
	}
}

//接收消息
func readMessage(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 409600)
	for {
		len := 0
		//读消息
		cnt, err := conn.Read(buf)
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
			fmt.Println("receive", conn.RemoteAddr(), stReceive)
			len += 12
		}
	}
}
