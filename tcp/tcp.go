package tcp

import (
	"fmt"
	"go-tcp-chat/services"
	"go-tcp-chat/utils"
	"net"
	"strings"
)

func HandleClient(conn net.Conn, a *int) {
	*a += 1
	b := *a
	fmt.Println("Handle connection %d", b)
	defer conn.Close()
	buffer := make([]byte, 1024) // Create a buffer to read data into

	for {
		// Read data from the client
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		buff_str := string(buffer)
		buff_parts := strings.Split(buff_str, " ")

		if !utils.IsRequestValid(buff_parts) {
			fmt.Fprintf(conn, "ERRO: faltou passar argumentos")
		}

		buff_parts[len(buff_parts)-1] = strings.ReplaceAll(buff_parts[len(buff_parts)-1], "\x00", "")
		requestType := strings.ToUpper(buff_parts[0])

		switch requestType {
		case "REGISTRO":
			user, _ := services.NewUser(buff_parts)
			msg, _ := services.HandleUserRegister(user)
			fmt.Fprintf(conn, msg)
		case "AUTENTICACAO":
			user, _ := services.NewUser(buff_parts)
			msg, _ := services.HandleUserAuthentication(user)
			fmt.Fprintf(conn, msg)
		case "CRIAR_SALA":
			fmt.Println("Criando sala %", buff_parts[1])

		default:
			fmt.Fprintf(conn, "ERRO: envie uma request válida\n")
		}

	}
}
