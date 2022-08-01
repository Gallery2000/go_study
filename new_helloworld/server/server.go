package main

import (
	"GoStudy/new_helloworld/handler"
	"GoStudy/new_helloworld/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	listener, _ := net.Listen("tcp", ":1234")
	_ = server_proxy.RegisterHelloService(&handler.NewHelloService{})
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
