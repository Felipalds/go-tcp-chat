package controllers

import (
	"crypto/rsa"
	"fmt"
	"go-tcp-chat/database"
	"go-tcp-chat/encrypt"
	"go-tcp-chat/models"
	"go-tcp-chat/services"
	"go-tcp-chat/utils"
	"net"
	"strings"
)

func HandleRequest(conn *net.Conn, buffParts []string, currentUser *models.User, pk **rsa.PrivateKey, aesKey *[]byte, auth *bool) (string, error) {
	requestType := strings.ToUpper(buffParts[0])
	var msg string
	var err error
	switch requestType {
	case "REGISTRO":
		user, err := services.NewUser(buffParts)

		msg, err = services.HandleUserRegister(user)
		if err != nil {
		}
		*currentUser = user
		msg = "REGISTRO_OK"
	case "AUTENTICACAO":
		userLoggin, _ := services.NewUser(buffParts)
		var userLogged models.User
		userLogged, msg, _ = services.HandleUserAuthentication(userLoggin)

		privateKey, _, err2 := encrypt.GenerateKeys()
		if err2 != nil {
			msg = "ERRO ao gerar chaves do usuário"
			fmt.Println(msg)
			break
		}

		*pk = privateKey
		NewClient(*conn, userLogged, privateKey)

		encodedKey, err2 := encrypt.EncodePublicToBase64(privateKey)
		if err2 != nil {
			fmt.Println(err2)
			break
		}
		msg = "CHAVE_PUBLICA " + encodedKey + "\n"
		*currentUser = userLogged

	case "CHAVE_SIMETRICA":
		aesKeyEncrypted := buffParts[1]
		*aesKey = encrypt.DecryptAESKey(aesKeyEncrypted, *pk)
		*auth = true
		msg = "AUTENTICACAO_OK\n"
	case "SAIR":
		*currentUser = models.User{}
		msg = utils.USER_LOGGED_OUT_MESSAGE

	case "CRIAR_SALA":
		fmt.Println("Criando sala")
		if !utils.IsLoggedIn(currentUser) {
			msg = utils.LOG_IN_FIRST_MESSAGE
			break
		}
		room, err2 := services.NewRoom(buffParts, *currentUser)
		if err2 != nil {
			fmt.Println(err2)
			break
		}
		msg, _ = services.HandleRoomRegister(room)

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
