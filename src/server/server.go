package server

import (
	"fmt"
	"log/slog"
	"net"
)

const defaultAddress string = "127.0.0.1:5050"

type ServerConfig struct {
	ListenAddress string
}

type Server struct {
	Config       ServerConfig
	Listener     net.Listener
	Peers        map[*Peer]bool
	AddPeerChann chan *Peer
	QuitChn      chan struct{}
}

func NewServer(serverCfg ServerConfig) *Server {
	if len(serverCfg.ListenAddress) == 0 {
		fmt.Println("[INFO]: Address not configured, Default address is being use")
		serverCfg.ListenAddress = defaultAddress
	}

	return &Server{
		Config:       serverCfg,
		Peers:        make(map[*Peer]bool),
		AddPeerChann: make(chan *Peer),
		QuitChn:      make(chan struct{}),
	}
}

func (s *Server) Start() error {

	slog.Info("[INFO][U-WORK]: STARTING SERVER", "listenAddress", s.Config.ListenAddress)

	ln, err := net.Listen("tcp", s.Config.ListenAddress)
	if err != nil {
		return err
	}
	s.Listener = ln

	//peer looper
	go s.Peerloop()

	//
	slog.Info("[INFO][SUCCESS]: SERVER STARTED ON: ", "listenAddress", s.Config.ListenAddress)

	if err := s.MainLoop(); err != nil {
		return err
	}

	return nil
}

func (s *Server) Peerloop() {
	for {
		select {
		case <-s.QuitChn:
			return
		case peer := <-s.AddPeerChann:
			s.Peers[peer] = true
		}
	}
}

func (s *Server) MainLoop() error {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			slog.Error("[ERROR][FATAL]: Error handling new connection", "err", err)
			continue
		}
		go s.ConnectionHandler(conn)
	}

}

func (s *Server) ConnectionHandler(conn net.Conn) {
	peer := NewPeer(conn)
	s.AddPeerChann <- peer
	slog.Info("[INFO] New peer Connected", "RemoteAddress", conn.RemoteAddr())
	go peer.MainLoop()
}
