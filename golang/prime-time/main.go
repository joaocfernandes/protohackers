package main

import (
	"encoding/json"
	"fmt"
	"net"
)

const PORT = ":8080"

type Payload struct {
	Method string `json:"method"`
	Number int    `json:"number"`
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)

	for {
		//read data from the client until it ends
		var payload Payload

		if err := decoder.Decode(&payload); err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		fmt.Println("Decoded payload:", payload)

		// Send a message back
		_, err := conn.Write([]byte("bla"))
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}

	}

}

func main() {
	// Start the server
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	defer listener.Close()

	fmt.Println("Server started.\n Listening on :", PORT)

	// Accept and handle incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		// For each new connection handle it in a separate goroutine
		go handleConnection(conn)
	}
}
