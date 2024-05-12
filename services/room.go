package services

import (
	"errors"
	"fmt"
	"go-tcp-chat/database"
	"go-tcp-chat/models"
	"go-tcp-chat/utils"
	"strings"
)

func NewRoom(params []string, user models.User) (models.Room, error) {
	// TODO: REFACTOR THIS FILE! AND CORRECT THE ERRORS
	var err error
	var password string

	if len(params) < 3 || len(params) > 4 {
		err = errors.New(utils.INVALID_ROOM_ARGUMENTS_MESSAGE)
	}

	roomType := strings.ToUpper(params[1])
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

func HandleRoomJoin(buffParts []string, user models.User) (string, error) {
	name := buffParts[1]
	var password string
	existedRoom, err := database.GetRoomByName(name)
	if existedRoom == nil || err != nil {
		return utils.ROOM_DOES_NOT_EXISTS_MESSAGE, err
	}
	if existedRoom.Type == "PRIVADA" {
		if len(buffParts) < 3 {
			return "PLEASE PROVIDE PASSWORD", nil
		}
		password = buffParts[2]
		if password != existedRoom.Password {
			return "INVALID PASSWORD", nil
		}
	}

	err2 := database.AddUserToRoom(user.Id, existedRoom.Id)
	if err2 != nil {
		return "ERROR ADDING USER TO ROOM\n", err2
	}
	return utils.USER_INSERTED_INTO_ROOM_MESSAGE, nil
}
