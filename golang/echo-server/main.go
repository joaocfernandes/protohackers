package main

import (
	"fmt"
	"net"
)

const PORT = ":8080"

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 2048)

	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		// Echo the received data back to the client
		_, err = conn.Write(buffer[:n])
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

	// Clean up when the server is stopped
	defer listener.Close()

	fmt.Println("Server started.\n Listening on :8080")

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
