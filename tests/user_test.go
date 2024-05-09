package tests

import (
	"go-tcp-chat/services"
	"testing"
)

func TestUserCreation(t *testing.T) {
	user_params := []string{"REGISTRO", "luiz"}
	_, err := services.HandleUserRegister(user_params)
	if err != nil {
		t.Errorf("Error while creating services: %v", err)
	}
}
