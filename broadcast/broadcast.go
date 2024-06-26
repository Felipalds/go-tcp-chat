package broadcast

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"go-tcp-chat/encrypt"
	"go-tcp-chat/models"
	"net"
	"sync"
)

type Client struct {
	conn  net.Conn
	user  models.User
	rooms []models.Room
	pk    *rsa.PrivateKey
	aes   []byte
}

var (
	clients   []*Client
	clientsMu sync.Mutex
)

func clientInGroup(client Client, roomName string) bool {
	for _, room := range client.rooms {
		if room.Name == roomName {
			return true
		}
	}
	return false
}

func InsertUserIntoRoom(user models.User, room models.Room) error {
	for _, client := range clients {
		if client.user.Name == user.Name {
			client.rooms = append(client.rooms, room)
			return nil
		}
	}
	return errors.New("Cannot append room to a user")
}

func NewClient(conn net.Conn, user models.User, pk *rsa.PrivateKey) {
	var newClient Client
	newClient.conn = conn
	newClient.user = user
	newClient.rooms = make([]models.Room, 0)
	newClient.pk = pk
	clients = append(clients, &newClient)
}

func UpdateClientAES(clientName string, aes []byte) {
	for _, client := range clients {
		if client.user.Name == clientName {
			client.aes = aes
		}
	}
}

func RemoveRoomFromClient(clientName string, roomName string) {
	for _, client := range clients {
		if client.user.Name == clientName {
			newRooms := make([]models.Room, 0)
			for i, room := range client.rooms {
				if room.Name == roomName {
					client.rooms = append(client.rooms[:i], client.rooms[i+1:]...)
				}
			}
			client.rooms = newRooms
		}
	}
}

func SendMessageToUser(userName string, msg string) error {
	for _, client := range clients {
		if client.user.Name == userName {
			msgEncrypted, err := encrypt.Encrypt([]byte(msg), client.aes)
			if err != nil {
				fmt.Println(err)
			}
			msgEncrypted += "\n"
			_, err = fmt.Fprintf(client.conn, msgEncrypted)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func UserIsInRoom(userName string, roomName string) bool {
	for _, client := range clients {
		if client.user.Name == userName {
			for _, room := range client.rooms {
				if room.Name == roomName {
					return true
				}
			}
		}
	}
	return false
}

func CloseRoom(roomName string) {
	for _, client := range clients {
		for i, room := range client.rooms {
			if room.Name == roomName {
				client.rooms = append(client.rooms[:i], client.rooms[i+1:]...)
			}
		}
		msg := "SALA_FECHADA " + roomName
		SendMessageToUser(client.user.Name, msg)
	}
}

func Broadcast(message string, roomName string, sender models.User) error {
	for _, client := range clients {
		if clientInGroup(*client, roomName) && sender.Name != client.user.Name {
			msgEncrypted, err := encrypt.Encrypt([]byte(message), client.aes)
			if err != nil {
				fmt.Println(err)
			}
			msgEncrypted += "\n"
			_, err = fmt.Fprintf(client.conn, msgEncrypted)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
