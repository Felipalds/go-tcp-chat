package controllers

import (
	"crypto/rsa"
	"fmt"
	"go-tcp-chat/encrypt"
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
	var aesKey []byte
	var auth bool
	auth = false

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
		buffStr = strings.Trim(buffStr, "\r\n")
		buffStr = strings.ReplaceAll(buffStr, "\x00", "")
		fmt.Println(len(buffStr))
		if auth {
			fmt.Println("Buff str", buffStr)
			buffStr, err = encrypt.Decrypt(buffStr, aesKey)
			fmt.Println("Buff str decytp", buffStr)
			if err != nil {
				fmt.Println("Error:", err)
				panic(err)
			}
			fmt.Println(buffStr)
		}

		buffParts := strings.Split(buffStr, " ")
		if !utils.IsRequestValid(buffParts) {
			fmt.Fprintf(conn, "ERRO: faltou passar argumentos")
		}

		buffParts[len(buffParts)-1] = strings.ReplaceAll(buffParts[len(buffParts)-1], "\x00", "")
		fmt.Println(buffParts)
		msg, _ := HandleRequest(&conn, buffParts, &user, &pk, &aesKey, &auth)
		fmt.Fprintf(conn, msg)
		buffer = make([]byte, 2048)
	}
}
