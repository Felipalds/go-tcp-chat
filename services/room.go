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
	var err error
	var password string
	var newRoom models.Room

	if len(params) < 3 || len(params) > 4 {
		return newRoom, errors.New(utils.INVALID_ROOM_ARGUMENTS_MESSAGE)
	}

	roomType := strings.ToUpper(params[1])
	name := strings.ReplaceAll(params[2], "\n", "")

	if roomType == "PRIVADA" && len(params) != 4 {
		return newRoom, errors.New(utils.ROOM_PASSWORD_NOT_PROVIDED_MESSAGE)
	}

	if len(params) == 4 {
		password = strings.ReplaceAll(params[3], "\n", "")
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

func HandleRoomJoin(room models.Room, user models.User) (string, error) {
	existedRoom, err := database.GetRoomByName(room.Name)
	if existedRoom == nil || err != nil {
		return utils.ROOM_DOES_NOT_EXISTS_MESSAGE, err
	}
	if existedRoom.Type == "PRIVADA" {
		if room.Password != existedRoom.Password {
			return "INVALID PASSWORD", nil
		}
	}

	err2 := database.AddUserToRoom(user.Id, existedRoom.Id)
	if err2 != nil {
		return "ERROR ADDING USER TO ROOM\n", err2
	}
	return utils.USER_INSERTED_INTO_ROOM_MESSAGE, nil
}

func GetRooms() (string, error) {
	msg := "SALAS "
	rooms, err := database.GetRooms()
	if err != nil {
		return "", err
	}
	for _, room := range rooms {
		msg += room.Name + " "
	}
	return msg, nil
}
