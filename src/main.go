package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "goServer/src/pb"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"syscall"
)
var(
	log = log.New()
)

func init(){
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	log.SetFormatter(&logrus.JSONFormatter{})
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	log.SetOutput(os.Stdout)
	//设置最低loglevel
	log.SetLevel(logrus.InfoLevel)
}

type Conn struct{
	conn *net.TCPConn
}
func (c *Conn) ok() bool{ return c!=nil && c.conn !=nil }

func(c *Conn) Write(b []byte)(int,error){
	if !c.ok(){
		return 0,syscall.EINVAL
	}
	n,err:=c.conn.Write(b)
	if err!=nil{

	}
	if n>0{

	}
}

func main() {
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	log.WithFields(logrus.Fields{
		"filename": "123.txt",
	}).Info("打开文件失败")

	listener, err := net.Listen("tcp", "localhost:6600")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
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
