package TcpServer

type TcpServer struct {
	TcpServer []*TcpConn
}

func CreateTcpServer() *TcpServer {
	t := &TcpServer{}
	t.TcpServer = make([]*TcpConn, 0)
	return t
}

func (c *TcpServer) Start() {

}

func (c *TcpServer) AddConn(conn *TcpConn) {
	if conn == nil {
		return
	}
	c.TcpServer = append(c.TcpServer, conn)
}

func (c *TcpServer) RemoveConnbyID(id uint64) {
	for k, v := range c.TcpServer {
		if v.serverid == id {
			c.TcpServer = append(c.TcpServer[:k], c.TcpServer[k+1:]...)
		}
	}
}

func (c *TcpServer) BroadCast(id uint64, str string) error {
	for _, v := range c.TcpServer {
		if v.serverid != id {
			v.conn.Write([]byte(str))
		}
	}
	return nil
}
