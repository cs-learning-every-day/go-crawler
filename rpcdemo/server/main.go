package main

import (
	"go-crawler/rpcdemo"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*
可使用telnet测试，一直黑屏就用ctrl + ]
telnet localhost 1234
{"method":"DemoService.Div", "parmas":[{"A":3,"B":4}], "id":1}
*/
func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
