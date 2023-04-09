package main

import (
	"fmt"
	"net"
)

type User struct {
	Name       string
	Address    string
	Channel    chan string
	Connection net.Conn
}

// NewUser 创建一个用户的接口
func NewUser(connection net.Conn) *User {
	userAddress := connection.RemoteAddr().String()

	user := User{
		Name:       userAddress,
		Address:    userAddress,
		Channel:    make(chan string),
		Connection: connection,
	}

	// 启动监听当前 User 的 Channel 消息的 goroutine
	go user.ListenMessage()

	return &user
}

// ListenMessage 监听当前 User 的 Channel 的 goroutine，一旦有消息就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		message := <-this.Channel

		_, err := this.Connection.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("this.Connection.Write error:", err)
			return
		}
	}
}
