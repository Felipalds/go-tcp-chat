package services

import (
	"errors"
	"fmt"
	"go-tcp-chat/database"
	"go-tcp-chat/models"
	"go-tcp-chat/utils"
)

func NewUser(params []string) (models.User, error) {
	var err error
	if len(params) != 2 {
		err = errors.New("User name not provided")
	}
	return models.User{Name: params[1]}, err
}

func HandleUserRegister(user models.User) (string, error) {
	MESSAGE := ""
	id, err := database.CreateNewUser(user.Name)
	if err != nil {
		fmt.Println(err)
	}
	if id == 0 {
		MESSAGE = utils.USER_ALREADY_EXISTS_MESSAGE
	} else {
		MESSAGE = utils.USER_REGISTERED_MESSAGE
	}
	return MESSAGE, nil
}

func HandleUserAuthentication(user models.User) (models.User, string, error) {
	MESSAGE := ""
	loggedUser, err := database.GetUserByName(user.Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loggedUser)
	if loggedUser.Name == "" {
		MESSAGE = utils.USER_DOES_NOT_EXISTS_MESSAGE
	} else {
		MESSAGE = utils.USER_LOGGED_MESSAGE
	}
	return *loggedUser, MESSAGE, nil
}
