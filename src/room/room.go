package room

type Room struct {
	room []*TcpConn
}

func CreateRoom() *Room{
	t:=&Room{}
	t.room=make([]*TcpConn,0)
	return t
}

func (c *Room)AddConn(conn *TcpConn){
	if conn==nil{
		return
	}
	c.room=append(c.room,conn)
}

func (c *Room)RemoveConnbyID(id uint64){
	for k,v:=range c.room{
		if v.connid==id{
			c.room=append(c.room[:k],c.room[k+1:]...)
		}
	}
}

func (c *Room)BroadCast(id uint64,str string) error{
	for _,v:=range c.room{
		if v.connid!=id{
			v.conn.Write([]byte(str))
		}
	}
}

