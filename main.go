package main

import (
	"fmt"
	"go-tcp-chat/controllers"
	"go-tcp-chat/database"
	"net"
)

func main() {
	db, _ := database.OpenDB()
	defer db.Close()
	database.Init(db)
	// listen for incomming connections
	fmt.Println("Starting server...")
	a := 0

	// to listen to connections until a message
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error creating server", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening!")

	for {
		// blocks code until data received
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error in listening connection", err)
			continue
		}

		// go routine here
		go controllers.HandleClient(conn, &a)
	}
}
