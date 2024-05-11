package tests

import (
	"go-tcp-chat/services"
	"go-tcp-chat/utils"
	"testing"
)

func TestUserCreation(t *testing.T) {
	user_params := []string{"REGISTRO", "luiz"}
	user, err := services.NewUser(user_params)

	if user.Name != "luiz" || err != nil {
		t.Errorf("User created incorrectly")
	}

	message, err := services.HandleUserRegister(user)
	if err != nil {
		t.Errorf("Error while creating services: %v", err)
	}
	if message != utils.USER_REGISTERED_MESSAGE {
		t.Errorf("Error while creating user, expected 'USER REGISTRED', got %s", message)
	}

	message2, err2 := services.HandleUserRegister(user)
	if err2 != nil {
		t.Errorf("Error while creating services: %v", err2)
	}
	if message2 != utils.USER_ALREADY_EXISTS_MESSAGE {
		t.Errorf("Error while creating user, expected 'ERROR : user already exists', got %s", message2)
	}

}
