package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	IP   string
	Port int

	OnlineUserMap  map[string]*User // 在线用户表
	mapLock        sync.RWMutex     // 用于在线用户表的读写锁
	MessageChannel chan string      // 用于消息广播的 Channel
}

// NewServer 创建一个 server 的接口
func NewServer(ip string, port int) *Server {
	server := Server{
		IP:             ip,
		Port:           port,
		OnlineUserMap:  make(map[string]*User),
		MessageChannel: make(chan string),
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

	// run MessageListener goroutine
	go this.MessageListener()

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
	isLive := make(chan bool) //监测当前负责的 User 是否仍然活跃

	// 用户上线
	// 1.创建新上线的用户
	user := NewUser(connection, this)

	// 2.将用户加入到在线用户表中，并广播当前用户上线的消息
	user.Online()

	// 3.监听 User 发送的消息
	go this.UserMessageListener(connection, user, &isLive)

	// 4.阻塞当前 Handler
	for {
		select {
		case <-isLive:
			// 当前负责的 User 仍然处于活跃状态，此时该条 case 触发，select 退出，进入下一次 for 循环，使得计时器被重置
		case <-time.After(time.Second * 20):
			// 已经超时，将当前负责的 User 强制关闭
			user.ShowMessage("Time out! offline!")
			close(user.Channel)                        // 销毁资源
			closeConnectionError := connection.Close() // 关闭连接
			if closeConnectionError != nil {
				fmt.Println("connection.Close error:", closeConnectionError)
			}
			return // 退出当前的 Handler
		}
	}
}

// Broadcast 广播当前用户上线的消息的方法
func (this *Server) Broadcast(user *User, message string) {
	broadcastMessage := "[" + user.Address + "]" + user.Name + ": " + message
	this.MessageChannel <- broadcastMessage
}

// MessageListener 监听 MessageChannel 的 goroutine，一旦有消息就马上发送给所有在线的 User
func (this *Server) MessageListener() {
	for {
		message := <-this.MessageChannel

		this.mapLock.Lock()
		for _, onlineUser := range this.OnlineUserMap {
			onlineUser.Channel <- message
		}
		this.mapLock.Unlock()
	}
}

// UserMessageListener 监听 User 发送的消息
func (this *Server) UserMessageListener(connection net.Conn, user *User, isLive *chan bool) {
	buffer := make([]byte, 4096)
	for {
		n, readError := connection.Read(buffer)
		if readError != nil && readError != io.EOF {
			fmt.Println("Connection.Read error:", readError)
			return
		}
		if n == 0 {
			user.Offline()
			return
		}

		message := string(buffer[:n-1]) // 提取用户的消息，并去除结尾的'\n'

		user.HandleMessage(message) // 用户对 message 进行处理

		*isLive <- true // 更新 User 的活跃状态
	}
}
