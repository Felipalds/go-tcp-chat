package controllers

import (
	"bufio"
	"crypto/rsa"
	"fmt"
	"go-tcp-chat/encrypt"
	"go-tcp-chat/models"
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

	fmt.Println("Handle a new connection %d", b)
	defer conn.Close()

	for {
		buffStr, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("A message received from client: ", buffStr)
		buffStr = strings.Trim(buffStr, "\r\n")
		buffStr = strings.ReplaceAll(buffStr, "\x00", "")
		if auth {
			buffStr, err = encrypt.Decrypt(buffStr, aesKey)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
		buffParts := strings.Split(buffStr, " ")

		buffParts[len(buffParts)-1] = strings.ReplaceAll(buffParts[len(buffParts)-1], "\x00", "")
		buffParts[len(buffParts)-1] = strings.ReplaceAll(buffParts[len(buffParts)-1], "\n", "")

		msg, err := HandleRequest(&conn, buffParts, &user, &pk, &aesKey, &auth)

		if err != nil {
			fmt.Println("Error:", err, err.Error())
			msg = "ERRO: " + err.Error()
		}

		if auth {
			var encryptErr error
			msg, encryptErr = encrypt.Encrypt([]byte(msg), aesKey)
			if encryptErr != nil {
				fmt.Println("Error encrypting msg to client:", err)
			}
		}

		msg += "\n"
		fmt.Println("MSG SENT: ", msg)
		fmt.Fprintf(conn, msg)
	}
}
