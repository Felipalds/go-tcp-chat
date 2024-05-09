package services

import (
	"errors"
	"fmt"
	"go-tcp-chat/database"
	"go-tcp-chat/models"
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
		MESSAGE = "ERROR : user already exists\n"
	} else {
		MESSAGE = "USER REGISTRED\n"
	}
	return MESSAGE, nil
}

func HandleUserAuthentication(user models.User) (string, error) {
	MESSAGE := ""
	loggedUser, err := database.GetUserByName(user.Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loggedUser)
	if loggedUser.Name == "" {
		MESSAGE = "ERROR : user does not exists\n"
	} else {
		MESSAGE = "USER LOGGED IN\n"
	}
	return MESSAGE, nil
}
