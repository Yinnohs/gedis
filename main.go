package main

import (
	"log"

	"github.com/yinnohs/gedis/src/server"
)

func main() {
	gedisServer := server.NewServer(server.ServerConfig{
		ListenAddress: "127.0.0.1:5050",
	})
	log.Fatal(gedisServer.Start())
}
