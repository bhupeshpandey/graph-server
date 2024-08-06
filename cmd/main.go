package main

import (
	"github.com/bhupeshpandey/graph-server/internal/server"
)

func main() {
	newServer := server.NewServer()
	err := newServer.Serve()
	if err != nil {
		return
	}
}
