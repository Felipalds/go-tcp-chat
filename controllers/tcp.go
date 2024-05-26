package controllers

import (
	"crypto/rsa"
	"fmt"
	"go-tcp-chat/models"
	"go-tcp-chat/utils"
	"net"
	"strings"
)

func HandleClient(conn net.Conn, a *int) {
	*a += 1
	b := *a
	var user models.User
	var pk *rsa.PrivateKey

	fmt.Println("Handle connection %d", b)
	defer conn.Close()
	buffer := make([]byte, 2048)

	for {
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		buffStr := string(buffer)
		buffParts := strings.Split(buffStr, " ")

		if !utils.IsRequestValid(buffParts) {
			fmt.Fprintf(conn, "ERRO: faltou passar argumentos")
		}

		buffParts[len(buffParts)-1] = strings.ReplaceAll(buffParts[len(buffParts)-1], "\x00", "")
		msg, _ := HandleRequest(&conn, buffParts, &user, &pk)
		fmt.Fprintf(conn, msg)
		buffer = make([]byte, 2048)
	}
}
