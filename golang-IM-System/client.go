package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	Connection net.Conn
}

func NewClient(serverIP string, serverPort int) *Client {
	// 创建 Client 对象
	client := Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
	}

	// 连接 Server
	connection, dialError := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if dialError != nil {
		fmt.Println("net.Dial error:", dialError)
		return nil
	}
	client.Connection = connection

	// 返回 Client 对象
	return &client
}

var serverIP string
var serverPort int

func init() {
	// ./client -ip 127.0.0.1 -port 8888
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "set server IP")
	flag.IntVar(&serverPort, "port", 8888, "set server port")
}

func main() {
	// 命令行解析
	flag.Parsed()

	client := NewClient(serverIP, serverPort)
	if client == nil {
		fmt.Println("Failed to connect server!")
		return
	}
	fmt.Println("Successfully connected to the server!")

	// 启动客户端的业务
	select {}
}
