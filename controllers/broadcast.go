package controllers

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"go-tcp-chat/models"
	"net"
	"sync"
)

type Client struct {
	conn  net.Conn
	user  models.User
	rooms []models.Room
	pk    *rsa.PrivateKey
}

var (
	clients   []*Client
	clientsMu sync.Mutex
)

func clientInGroup(client Client, roomName string) bool {
	for _, room := range client.rooms {
		fmt.Println(room.Name, roomName)
		if room.Name == roomName {
			return true
		}
	}
	return false
}

func InsertUserIntoRoom(user models.User, room models.Room) error {
	fmt.Println("HERE2")
	for _, client := range clients {
		if client.user.Name == user.Name {
			client.rooms = append(client.rooms, room)
			return nil
		}
	}
	return errors.New("Cannot append room to a user")
}

func NewClient(conn net.Conn, user models.User, pk *rsa.PrivateKey) {
	clientsMu.Lock()
	var newClient Client
	newClient.conn = conn
	newClient.user = user
	newClient.rooms = make([]models.Room, 0)
	newClient.pk = pk
	clients = append(clients, &newClient)
	clientsMu.Unlock()
}

func Broadcast(message []string, roomName string, sender models.User) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for _, client := range clients {
		if clientInGroup(*client, roomName) && sender.Name != client.user.Name {
			fmt.Fprintf(client.conn, "%s\n", message)
		}
	}
}
