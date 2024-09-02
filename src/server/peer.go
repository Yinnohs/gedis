package server

import (
	"net"
)

type Peer struct {
	conn net.Conn
}

func NewPeer(conn net.Conn) *Peer {
	return &Peer{
		conn: conn,
	}
}

func (p *Peer) MainLoop() error {
	buff := make([]byte, 1024)
	for {
		dataBytes, err := p.conn.Read(buff)
		if err != nil {
			return err
		}
		msgBuffer := make([]byte, dataBytes)
		copy(msgBuffer, buff[:dataBytes])

	}
}
