package main

import (
	"net"
	"sync/atomic"
	"time"
)

const (
	kServerConf_SendBufferSize = 1024
	kServerConn                = 0
	kClientConn                = 1
)

// TCPNetworkConf config the TCPNetwork
type TCPNetworkConf struct {
	SendBufferSize int
}

// TCPNetwork manages all server and client connections
type TCPNetwork struct {
	streamProtocol  IStreamProtocol
	eventQueue      chan *ConnEvent
	listener        net.Listener
	Conf            TCPNetworkConf
	connIdForServer int
	connIdForClient int
	connsForServer  map[int]*Connection
	connsForClient  map[int]*Connection
	shutdownFlag    int32
	readTimeoutSec  int
}

// NewTCPNetwork creates a TCPNetwork object
func NewTCPNetwork(eventQueueSize int, sp IStreamProtocol) *TCPNetwork {
	s := &TCPNetwork{}
	s.eventQueue = make(chan *ConnEvent, eventQueueSize)
	s.streamProtocol = sp
	s.connsForServer = make(map[int]*Connection)
	s.connsForClient = make(map[int]*Connection)
	s.shutdownFlag = 0
	//	default config
	s.Conf.SendBufferSize = kServerConf_SendBufferSize
	return s
}

// Push implements the IEventQueue interface
func (t *TCPNetwork) Push(evt *ConnEvent) {
	if nil == t.eventQueue {
		return
	}

	//	push timeout
	select {
	case t.eventQueue <- evt:
		{

		}
	case <-time.After(time.Second * 5):
		{
			evt.Conn.close()
		}
	}

}

// Pop the event in event queue
func (t *TCPNetwork) Pop() *ConnEvent {
	evt, ok := <-t.eventQueue
	if !ok {
		//	event queue already closed
		return nil
	}

	return evt
}

// GetEventQueue get the event queue channel
func (t *TCPNetwork) GetEventQueue() <-chan *ConnEvent {
	return t.eventQueue
}

// Listen an address to accept client connection
func (t *TCPNetwork) Listen(addr string) error {
	ls, err := net.Listen("tcp", addr)
	if nil != err {
		return err
	}

	//	accept
	t.listener = ls
	go t.acceptRoutine()
	return nil
}

// Connect the remote server
func (t *TCPNetwork) Connect(addr string) (*Connection, error) {
	conn, err := net.Dial("tcp", addr)
	if nil != err {
		return nil, err
	}

	connection := t.createConn(conn)
	connection.from = 1
	connection.run()
	connection.init()

	return connection, nil
}

// GetStreamProtocol returns the stream protocol of current TCPNetwork
func (t *TCPNetwork) GetStreamProtocol() IStreamProtocol {
	return t.streamProtocol
}

// SetStreamProtocol sets the stream protocol of current TCPNetwork
func (t *TCPNetwork) SetStreamProtocol(sp IStreamProtocol) {
	t.streamProtocol = sp
}

// GetReadTimeoutSec returns the read timeout seconds of current TCPNetwork
func (t *TCPNetwork) GetReadTimeoutSec() int {
	return t.readTimeoutSec
}

// SetReadTimeoutSec sets the read timeout seconds of current TCPNetwork
func (t *TCPNetwork) SetReadTimeoutSec(sec int) {
	t.readTimeoutSec = sec
}

// DisconnectAllConnectionsServer disconnect all connections on server side
func (t *TCPNetwork) DisconnectAllConnectionsServer() {
	for k, c := range t.connsForServer {
		c.Close()
		delete(t.connsForServer, k)
	}
}

// Shutdown frees all connection and stop the listener
func (t *TCPNetwork) Shutdown() {
	if !atomic.CompareAndSwapInt32(&t.shutdownFlag, 0, 1) {
		return
	}

	// Stop accept routine
	if nil != t.listener {
		t.listener.Close()
	}

	// Close all connections
	t.DisconnectAllConnectionsClient()
	t.DisconnectAllConnectionsServer()
}

func (t *TCPNetwork) createConn(c net.Conn) *Connection {
	conn := newConnection(c, t.Conf.SendBufferSize, t)
	conn.setStreamProtocol(t.streamProtocol)
	return conn
}
