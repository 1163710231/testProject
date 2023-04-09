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
	Flag       int
}

func NewClient(serverIP string, serverPort int) *Client {
	// 创建 Client 对象
	client := Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
		Flag:       999,
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

func (this *Client) showMenu() bool {
	fmt.Println("0.退出")
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")

	var f int
	_, ScanError := fmt.Scanln(&f)
	if ScanError != nil {
		fmt.Println("fmt.Scan Error:", ScanError)
		return false
	}
	if 0 <= f && f <= 3 {
		this.Flag = f
		return true
	} else {
		fmt.Println("Input Error!")
		return false
	}
}

func (this *Client) Run() {
	for this.Flag != 0 {
		for this.showMenu() == false {

		}

		// 根据不同的业务模式处理不同的业务
		switch this.Flag {
		case 0: // 退出
			fmt.Println("退出")
			break
		case 1: // 公聊模式
			fmt.Println("公聊模式")
			break
		case 2: // 私聊模式
			fmt.Println("私聊模式")
			break
		case 3: // 更新用户名
			fmt.Println("更新用户名")
			break
		default: // 错误
			fmt.Println("Model error!")
		}
	}
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
	client.Run()
}
