package tests

import (
	"fmt"
	"go-tcp-chat/controllers"
	"go-tcp-chat/models"
	"go-tcp-chat/utils"
	"testing"
)

func TestController(t *testing.T) {

	var user models.User
	buffParts := []string{"CRIAR_SALA", "PUBLICA", "LUIZ"}
	msg, err := controllers.HandleRequest(buffParts, &user)
	if err != nil || msg != utils.LOG_IN_FIRST_MESSAGE {
		t.Error("ERROR HANDLING WITH CREATING ROOM WITHOUT LOGIN", err)
	}

	buffParts = []string{"REGISTRO", "USER123"}
	msg, err = controllers.HandleRequest(buffParts, &user)
	if err != nil || msg != utils.USER_REGISTERED_MESSAGE {
		t.Error("ERROR HANDLING WITH REGISTER", err)
	}

	buffParts = []string{"CRIAR_SALA", "PUBLICA", "LUIZ"}
	msg, err = controllers.HandleRequest(buffParts, &user)
	if err != nil || msg != utils.LOG_IN_FIRST_MESSAGE {
		t.Error("ERROR HANDLING WITH CREATING ROOM WITHOUT LOGIN", err)
	}

	buffParts = []string{"AUTENTICACAO", "USER123"}
	msg, err = controllers.HandleRequest(buffParts, &user)
	if err != nil || msg != utils.USER_LOGGED_MESSAGE {
		t.Error("ERROR HANDLING WITH LOGGED", err)
	}

	buffParts = []string{"CRIAR_SALA", "PUBLICA", "LUIZ"}
	msg, err = controllers.HandleRequest(buffParts, &user)
	if err != nil || msg != utils.ROOM_CREATED_SUCCESS_MESSAGE {
		fmt.Println(msg)
		t.Error("ERROR HANDLING WITH CREATING ROOM AFTER LOGIN", err)
	}

	buffParts = []string{"CRIAR_SALA", "PUBLICA", "LUIZ"}
	msg, err = controllers.HandleRequest(buffParts, &user)
	if err != nil || msg != utils.ROOM_ALREADY_EXISTS_MESSAGE {
		fmt.Println(msg)
		t.Error("ERROR HANDLING WITH CREATING ROOM ALREADY EXISTS", err)
	}

	//buffParts = []string{"CRIAR_SALA", "PRIVADA", "TESTE2"}
	//msg, err = controllers.HandleRequest(buffParts, &user)
	//if err != nil && msg != utils.ROOM_PASSWORD_NOT_PROVIDED_MESSAGE {
	//	fmt.Println(msg)
	//	t.Error("ERROR HANDLING WITH ROOM WITHOUT PASSWORD", err)
	//}

	//buffParts = []string{"CRIAR_SALA", "PRIVADA", "TESTE2", "abc"}
	//msg, err = controllers.HandleRequest(buffParts, &user)
	//if err != nil || msg != utils.ROOM_CREATED_SUCCESS_MESSAGE {
	//	t.Error("ERROR HANDLING WITH ROOM WITHOUT PASSWORD", err)
	//}

	buffParts = []string{"REGISTRO", "USER123"}
	msg, err = controllers.HandleRequest(buffParts, &user)
	if err != nil || msg != utils.USER_ALREADY_EXISTS_MESSAGE {
		t.Error("ERROR HANDLING WITH EXISTING USER", err)
	}

}
