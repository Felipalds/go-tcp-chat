package tcp

import (
	"fmt"
	"go-tcp-chat/user"
	"net"
	"strings"
)

func HandleClient(conn net.Conn, a *int) {
	*a += 1
	b := *a
	fmt.Println("Handle connection ", *a)
	defer conn.Close()
	buffer := make([]byte, 1024) // Create a buffer to read data into

	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		buff_str := string(buffer)
		buff_parts := strings.Split(buff_str, " ")

		if len(buff_parts) == 0 {
			fmt.Fprintf(conn, "ERRO: faltou passar argumentos")
			return
		}

		if buff_parts[0] == "REGISTRO" {
			msg, err := user.HandleUserRegister(buff_parts)
			if err != nil {
				fmt.Fprintf(conn, msg)
			}
			fmt.Fprintf(conn, msg)
		}
		//
		//if buff_parts[0] == "CRIAR_SALA" {
		//	if len(buff_parts) < 4 {
		//		fmt.Fprintf(conn, "ERRO: faltou passar argumento")
		//		return
		//	}
		//	fmt.Println("Criando sala %", buff_parts[1])
		//}

		// Process and use the data (here, we'll just print it)
		fmt.Printf("Received %d %d bytes: %s\n", b, n, buffer[:n])
	}
}
