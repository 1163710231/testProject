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

	ConnectedServers *Server
}

// NewUser 创建一个用户的接口
func NewUser(connection net.Conn, connectedServers *Server) *User {
	userAddress := connection.RemoteAddr().String()

	user := User{
		Name:             userAddress,
		Address:          userAddress,
		Channel:          make(chan string),
		Connection:       connection,
		ConnectedServers: connectedServers,
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

// Online 用户的上线功能
func (this *User) Online() {
	// 将用户加入到在线用户表中
	this.ConnectedServers.mapLock.Lock()
	this.ConnectedServers.OnlineUserMap[this.Name] = this
	this.ConnectedServers.mapLock.Unlock()

	// 广播当前用户上线的消息
	this.ConnectedServers.Broadcast(this, "is online now!")
}

// Offline 用户的下线功能
func (this *User) Offline() {
	// 将用户从在线用户表中删除
	this.ConnectedServers.mapLock.Lock()
	delete(this.ConnectedServers.OnlineUserMap, this.Name)
	this.ConnectedServers.mapLock.Unlock()

	// 广播当前用户下线的消息
	this.ConnectedServers.Broadcast(this, "is offline now!")
}

// HandleMessage 用户的处理消息功能
func (this *User) HandleMessage(message string) {
	this.ConnectedServers.Broadcast(this, message)
}
