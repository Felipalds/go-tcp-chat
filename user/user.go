package user

import (
	"errors"
	"fmt"
)

func HandleUserRegister(params []string) (string, error) {
	fmt.Println("User register")

	if len(params) != 2 {
		ERROR_MESSAGE := "USER WITHOUT NAME"
		fmt.Println(ERROR_MESSAGE)
		return ERROR_MESSAGE, errors.New("USER WITHOUT NAME")
	}

	// register user
	SUCCESS_MESSAGE := "USER REGISTER"
	return SUCCESS_MESSAGE, nil
}
