package services

import (
	"errors"
	"fmt"
	"go-tcp-chat/database"
	"go-tcp-chat/models"
	"go-tcp-chat/utils"
)

func NewRoom(params []string, user models.User) (models.Room, error) {
	// TODO: REFACTOR THIS FILE! AND CORRECT THE ERRORS
	var err error
	var password string

	if len(params) < 3 || len(params) > 4 {
		err = errors.New(utils.INVALID_ROOM_ARGUMENTS_MESSAGE)
	}

	roomType := params[1]
	name := params[2]

	if roomType != "PUBLICA" && len(params) != 4 {
		err = errors.New(utils.ROOM_PASSWORD_NOT_PROVIDED_MESSAGE)
	}

	if len(params) == 4 {
		password = params[3]
	}

	return models.Room{Name: name, Type: roomType, Password: password, Admin: user}, err
}

func HandleRoomRegister(room models.Room) (string, error) {
	MESSAGE := ""
	id, err := database.CreateNewRoom(room)
	if err != nil {
		fmt.Println(err)
	}
	if id == 0 {
		MESSAGE = utils.ROOM_ALREADY_EXISTS_MESSAGE
	} else {
		MESSAGE = utils.ROOM_CREATED_SUCCESS_MESSAGE
	}
	return MESSAGE, nil
}

//func HandleUserAuthentication(user models.User) (string, error) {
//	MESSAGE := ""
//	loggedUser, err := database.GetUserByName(user.Name)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(loggedUser)
//	if loggedUser.Name == "" {
//		MESSAGE = "ERROR : user does not exists\n"
//	} else {
//		MESSAGE = "USER LOGGED IN\n"
//	}
//	return MESSAGE, nil
//}
