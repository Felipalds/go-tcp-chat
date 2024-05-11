package controllers

import (
	"fmt"
	"go-tcp-chat/models"
	"go-tcp-chat/services"
	"go-tcp-chat/utils"
	"net"
	"strings"
)

func HandleClient(conn net.Conn, a *int) {
	*a += 1
	b := *a
	var user models.User
	fmt.Println("Handle connection %d", b)
	defer conn.Close()
	buffer := make([]byte, 1024) // Create a buffer to read data into

	for {
		// Read data from the client
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		buffStr := string(buffer)
		buffParts := strings.Split(buffStr, " ")

		if !utils.IsRequestValid(buffParts) {
			fmt.Fprintf(conn, "ERRO: faltou passar argumentos")
		}

		buffParts[len(buffParts)-1] = strings.ReplaceAll(buffParts[len(buffParts)-1], "\x00", "")
		msg, _ := HandleRequest(buffParts, &user)
		fmt.Fprintf(conn, msg)
		buffer = make([]byte, 1024)

	}
}

func HandleRequest(buffParts []string, currentUser *models.User) (string, error) {
	requestType := strings.ToUpper(buffParts[0])
	var msg string
	var err error
	switch requestType {
	case "REGISTRO":
		user, _ := services.NewUser(buffParts)
		msg, err = services.HandleUserRegister(user)
	case "AUTENTICACAO":
		userLoggin, _ := services.NewUser(buffParts)
		var userLogged models.User
		userLogged, msg, _ = services.HandleUserAuthentication(userLoggin)
		*currentUser = userLogged

	case "SAIR":
		*currentUser = models.User{}
		msg = utils.USER_LOGGED_OUT_MESSAGE

	case "CRIAR_SALA":
		if currentUser == nil || currentUser.Name == "" {
			msg = utils.LOG_IN_FIRST_MESSAGE
			break
		}
		room, _ := services.NewRoom(buffParts, *currentUser)
		msg, err = services.HandleRoomRegister(room)

	default:
		msg = "REQUEST INVÁLIDA"
	}

	return msg, err
}
