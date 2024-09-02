package server

import (
	"net"
)

type Peer struct {
	conn    net.Conn
	MsgChan chan []byte
}

func NewPeer(conn net.Conn, msgChan chan []byte) *Peer {
	return &Peer{
		conn:    conn,
		MsgChan: msgChan,
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
		p.MsgChan <- msgBuffer
	}
}
