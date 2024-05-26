package tests

import (
	"fmt"
	"go-tcp-chat/models"
	"go-tcp-chat/services"
	"testing"
)

func TestRoomCreation(t *testing.T) {

	name := "teste2\n"
	password := "teste1"
	user := models.User{
		Id:   0,
		Name: "luizinho",
	}
	buffParts := []string{"CRIAR_SALA", "PRIVADA", name, password}

	fmt.Println("Criando sala...")
	room, err := services.NewRoom(buffParts, user)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	_, err = services.HandleRoomRegister(room)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

}
