package controllers

import (
	"crypto/rsa"
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

func InsertUserIntoRoom(conn net.Conn, user models.User, room models.Room) {
	var existingClient *Client
	for _, client := range clients {
		if client.user.Name == user.Name {
			existingClient.rooms = append(existingClient.rooms, room)
			clients = append(clients, existingClient)
			return
		}
	}

	var newClient = &Client{conn: conn, user: user, rooms: []models.Room{}}
	newClient.rooms = append(newClient.rooms, room)
	clients = append(clients, newClient)
}

func NewClient(conn net.Conn, user models.User, pk *rsa.PrivateKey) {
	clientsMu.Lock()
	var newClient Client
	newClient.conn = conn
	newClient.user = user
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
