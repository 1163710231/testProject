package main

import (
	"fmt"
	"net"
	"strings"
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

// ListenMessage 监听当前 User 的 Channel 的 goroutine，一旦有消息就直接进行发送给本 User，以供其显示
func (this *User) ListenMessage() {
	for {
		message := <-this.Channel

		this.ShowMessage(message)
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
	if message == "who" { // 查询当前所有的在线用户
		this.ConnectedServers.mapLock.Lock()
		for _, user := range this.ConnectedServers.OnlineUserMap {
			onlineCountMessage := "[" + user.Address + "]" + user.Name + ": is online now..."
			this.ShowMessage(onlineCountMessage)
		}
		this.ConnectedServers.mapLock.Unlock()
	} else if len(message) > 7 && message[:7] == "rename|" {
		newName := strings.Split(message, "|")[1]
		_, ok := this.ConnectedServers.OnlineUserMap[newName]
		if ok { // 用户名已经存在
			this.ShowMessage("User name [" + newName + "] already exist\n")
		} else {
			this.ConnectedServers.mapLock.Lock()
			delete(this.ConnectedServers.OnlineUserMap, this.Name)
			this.ConnectedServers.OnlineUserMap[newName] = this
			this.ConnectedServers.mapLock.Unlock()

			this.Name = newName

			this.ShowMessage("User name updated: " + newName + "\n")
		}
	} else {
		this.ConnectedServers.Broadcast(this, message)
	}
}

// ShowMessage 将消息发送给本用户，从而在其界面显示消息
func (this *User) ShowMessage(message string) {
	_, err := this.Connection.Write([]byte(message + "\n"))
	if err != nil {
		fmt.Println("this.Connection.Write error:", err)
		return
	}
}
