package main

import (
	"net"
	"time"
)


func main(){
	conn, err := net.DialTimeout("tcp", "localhost:6600", 2 * time.Second)
	if err != nil {
		//handle error
	}
}
