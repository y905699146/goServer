package main

import (
	"fmt"
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
		go pl.ReadMessage()
	}
}

