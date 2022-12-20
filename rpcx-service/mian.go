package main

import (
	"github.com/smallnest/rpcx/server"
	"rpcx-service/service"
)

func main() {
	server := server.NewServer()
	server.RegisterName("Arith", new(service.Arith), "")
	server.Serve("tcp", ":8972")
}
