package main

import (
	"fmt"
	"goServer/src/logger"
	"goServer/src/room"
	"net"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	logger.Debug_MSG("server close down(signal: %v)", sig)
	listener, err := net.Listen("tcp", "localhost:6600")
	if err != nil {
		panic(err)
	}
	r := room.CreateRoom()
	for {
		conn, err := listener.Accept()
		pl := room.CreateConn(conn)
		r.AddConn(pl)
		if err != nil {
			panic(err)
		}
		fmt.Println("new connect", conn.RemoteAddr())
		go pl.ReadMessage()
	}
}
