package controllers

import (
	"crypto/rsa"
	"fmt"
	"go-tcp-chat/encrypt"
	"go-tcp-chat/models"
	"go-tcp-chat/services"
	"go-tcp-chat/utils"
	"net"
	"strings"
)

func HandleRequest(conn *net.Conn, buffParts []string, currentUser *models.User, pk **rsa.PrivateKey, aesKey *[]byte, auth *bool) (string, error) {
	requestType := strings.ToUpper(buffParts[0])
	requestType = strings.ReplaceAll(requestType, "\n", "")
	fmt.Println("Logs for request: ", requestType) //dont remove
	var msg string
	var err error
	switch requestType {
	case "REGISTRO":
		user, err := services.NewUser(buffParts)
		if err != nil {
			return "", err
		}

		msg, err = services.HandleUserRegister(user)
		if err != nil {
			return "", err
		}
		msg = "REGISTRO_OK\n"
	case "AUTENTICACAO":
		userLoggin, err := services.NewUser(buffParts)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		userLogged, err := services.HandleUserAuthentication(userLoggin)
		if err != nil {
			fmt.Println(err)
			return "", err
		}

		privateKey, _, err2 := encrypt.GenerateKeys()
		if err2 != nil {
			msg = "ERRO ao gerar chaves do usuário"
			fmt.Println(msg)
			break
		}

		*pk = privateKey
		NewClient(*conn, *userLogged, privateKey)

		encodedKey, err2 := encrypt.EncodePublicToBase64(privateKey)
		if err2 != nil {
			fmt.Println(err2)
			break
		}

		msg = "CHAVE_PUBLICA " + encodedKey
		*currentUser = *userLogged
		fmt.Println(currentUser)

	case "CHAVE_SIMETRICA":
		aesKeyEncrypted := buffParts[1]
		*aesKey = encrypt.DecryptAESKey(aesKeyEncrypted, *pk)
		*auth = true
		msg = "AUTENTICACAO_OK"
	case "SAIR":
		*currentUser = models.User{}
		msg = utils.USER_LOGGED_OUT_MESSAGE

	case "CRIAR_SALA":
		fmt.Println("Criando sala...")
		if !utils.IsLoggedIn(currentUser) {
			return utils.LOG_IN_FIRST_MESSAGE, nil
		}
		room, err := services.NewRoom(buffParts, *currentUser)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		msg, err = services.HandleRoomRegister(room)
		if err != nil {
			fmt.Println(err)
			return "", err
		}

		err = InsertUserIntoRoom(*currentUser, room)
		if err != nil {
			return "", err
		}

		return utils.ROOM_CREATED_SUCCESS_MESSAGE, nil

	case "ENTRAR_SALA":
		if !utils.IsLoggedIn(currentUser) {
			return utils.LOG_IN_FIRST_MESSAGE, nil
		}
		room, err := services.HandleRoomJoin(buffParts, *currentUser)
		if err != nil {
			return "", err
		}

		err = InsertUserIntoRoom(*currentUser, room)
		if err != nil {
			return "", err
		}
		return "ENTRAR_SALA_OK", nil

	case "LISTAR_SALAS":
		if !utils.IsLoggedIn(currentUser) {
			return utils.LOG_IN_FIRST_MESSAGE, nil
		}
		rooms, err := services.GetRooms()
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		return rooms, nil

	case "ENVIAR_MENSAGEM":
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
