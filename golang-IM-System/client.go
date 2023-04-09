package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
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

func (this *Client) ReceiveMessageFromServer() {
	_, ioCopyError := io.Copy(os.Stdout, this.Connection) // 一旦 this.Connection 中有输出，就直接 copy 到标准输出上，且永久阻塞监听
	if ioCopyError != nil {
		fmt.Println("io.Copy error:", ioCopyError)
		return
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
			this.UpdateName()
			break
		default: // 错误
			fmt.Println("Model error!")
		}
	}
}

func (this *Client) UpdateName() bool {
	fmt.Println("Please input new user name")
	_, ScanError := fmt.Scanln(&this.Name)
	if ScanError != nil {
		fmt.Println("Scan error:", ScanError)
		return false
	}

	renameMessage := "rename|" + this.Name + "\n"
	_, connectionWriteError := this.Connection.Write([]byte(renameMessage))
	if connectionWriteError != nil {
		fmt.Println("Connection.Write error:", connectionWriteError)
		return false
	}
	return true
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

	go client.ReceiveMessageFromServer()

	fmt.Println("Successfully connected to the server!")

	// 启动客户端的业务
	client.Run()
}
