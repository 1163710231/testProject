package main

import (
	"fmt"
	"net"
)

type Server struct {
	IP   string
	Port int
}

// NewServer 创建一个 server 的接口
func NewServer(ip string, port int) *Server {
	server := Server{
		IP:   ip,
		Port: port,
	}
	return &server
}

// Start 启动服务器的接口
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.Port))
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}

	// defer close listen socket
	defer func(listener net.Listener) {
		closeError := listener.Close()
		if closeError != nil {
			fmt.Println("listener.Close error:", closeError)
		}
	}(listener)

	for {
		// accept
		connection, connectionError := listener.Accept()
		if connectionError != nil {
			fmt.Println("listener.Accept error:", connectionError)
			continue
		}

		// do handler
		go this.Handler(connection)
	}
}

func (this *Server) Handler(connection net.Conn) {
	// 当前连接的业务
	fmt.Println("执行当前连接的业务")
}
