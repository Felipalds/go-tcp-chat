package controllers

import (
	"crypto/rsa"
	"fmt"
	"go-tcp-chat/database"
	"go-tcp-chat/models"
	"go-tcp-chat/services"
	"go-tcp-chat/utils"
	"net"
	"strings"
)

func HandleRequest(conn *net.Conn, buffParts []string, currentUser *models.User, pk **rsa.PrivateKey) (string, error) {
	requestType := strings.ToUpper(buffParts[0])
	var msg string
	var err error
	switch requestType {
	case "REGISTRO":
		user, _ := services.NewUser(buffParts)
		msg, err = services.HandleUserRegister(user)
		if err != nil {
			*currentUser = user
		}
	case "AUTENTICACAO":
		userLoggin, _ := services.NewUser(buffParts)
		var userLogged models.User
		userLogged, msg, _ = services.HandleUserAuthentication(userLoggin)

		privateKey, _, err2 := GenerateKeys()
		if err2 != nil {
			msg = "ERRO ao gerar chaves do usuário"
			fmt.Println(msg)
			break
		}

		*pk = privateKey
		fmt.Println("Public Key", (*pk).PublicKey)
		NewClient(*conn, userLogged, privateKey)

		encodedKey, err2 := EncodePublicToBase64(privateKey)
		if err2 != nil {
			fmt.Println(err2)
			break
		}
		fmt.Println("PUBLIC ENCODED", encodedKey)
		msg = "CHAVE_PUBLICA " + encodedKey + "\n"
		*currentUser = userLogged

	case "CHAVE_SIMETRICA":
		aesKey := buffParts[1]
		DecryptAESKey(aesKey, *pk)
	case "SAIR":
		*currentUser = models.User{}
		msg = utils.USER_LOGGED_OUT_MESSAGE

	case "CRIAR_SALA":
		if !utils.IsLoggedIn(currentUser) {
			msg = utils.LOG_IN_FIRST_MESSAGE
			break
		}
		room, _ := services.NewRoom(buffParts, *currentUser)
		msg, err = services.HandleRoomRegister(room)
		InsertUserIntoRoom(*conn, *currentUser, room)

	case "ENTRAR_SALA":
		if !utils.IsLoggedIn(currentUser) {
			msg = utils.LOG_IN_FIRST_MESSAGE
			break
		}
		// TODO: achar uma maneira melhor de fazer os replace all e de lidar com os buffParts
		room, _ := database.GetRoomByName(strings.ReplaceAll(buffParts[1], "\n", ""))
		msg, _ = services.HandleRoomJoin(*room, *currentUser)
		InsertUserIntoRoom(*conn, *currentUser, *room)

	case "ENVIAR_MSG":
		if !utils.IsLoggedIn(currentUser) {
			msg = utils.LOG_IN_FIRST_MESSAGE
			break
		}
		roomName := strings.ReplaceAll(buffParts[1], " ", "")
		Broadcast(buffParts[2:], roomName, *currentUser)
	default:
		msg = "REQUEST INVÁLIDA"
	}

	return msg, err
}
