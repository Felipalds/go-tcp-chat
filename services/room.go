package services

import (
	"errors"
	"fmt"
	"go-tcp-chat/broadcast"
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

func HandleRoomJoin(buffParts []string, user models.User) (models.Room, error) {

	roomName := ""
	roomPassword := ""
	var room models.Room

	if len(buffParts) < 2 || len(buffParts) > 3 {
		return room, errors.New("invalid")
	}

	if len(buffParts) == 3 {
		roomPassword = buffParts[2]
	}

	roomName = buffParts[1]
	roomName = strings.ReplaceAll(roomName, "\n", "")

	existedRoom, err := database.GetRoomByName(roomName)
	if existedRoom == nil || err != nil {
		return room, errors.New(utils.ROOM_DOES_NOT_EXISTS_MESSAGE)
	}
	room = *existedRoom
	if existedRoom.Type == "PRIVADA" {
		fmt.Println(existedRoom.Password, roomPassword)
		if existedRoom.Password != roomPassword {

			return room, errors.New("Invalid password")
		}
	}

	err2 := database.AddUserToRoom(user.Id, existedRoom.Id)
	if err2 != nil {
		return room, errors.New("Error adding user to room")
	}
	return room, nil
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

func HandleBan(buffParts []string, user models.User) error {

	var userToBeBanned *models.User

	if len(buffParts) < 3 {
		return errors.New("not enough arguments to ban a user")
	}

	room, err := database.GetRoomByName(buffParts[1])
	if err != nil {
		return err
	}

	if room.Admin.Name != user.Name {
		return errors.New("YOU ARE NOT ADMIN")
	}

	userToBeBanned, err = database.GetUserByName(buffParts[2])
	if err != nil {
		return err
	}

	broadcast.RemoveRoomFromClient(userToBeBanned.Name, room.Name)

	return nil
}

func HandleLeave(buffParts []string, user models.User) error {

	if len(buffParts) < 2 {
		return errors.New("not enough arguments to leave a user")
	}

	room, err := database.GetRoomByName(buffParts[1])
	if err != nil {
		return err
	}

	if room.Admin.Name != user.Name {
		broadcast.CloseRoom(room.Name)
	} else {
		broadcast.RemoveRoomFromClient(user.Name, room.Name)
	}

	return nil

}

func HandleCloseRoom(buffParts []string, user models.User) error {

	if len(buffParts) < 2 {
		return errors.New("not enough arguments to ban a user")
	}

	room, err := database.GetRoomByName(buffParts[1])
	if err != nil {
		return err
	}

	if room.Admin.Name != user.Name {
		return errors.New("YOU ARE NOT ADMIN")
	}
	broadcast.CloseRoom(buffParts[1])
	return nil
}
