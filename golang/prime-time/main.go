// https://protohackers.com/problem/1

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"

	"github.com/fxtlabs/primes"
)

const PORT = ":8080"

type Request struct {
	Method string `json:"method"`
	Number int    `json:"number"`
}

type Response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

func responseOnError(conn net.Conn) {
	// Send a error message back
	_, err := conn.Write([]byte("ERROR"))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
	}
	return
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	//read data from the client until it ends
	for scanner.Scan() {
		var payload Request

		line := scanner.Text()

		if err := json.Unmarshal([]byte(line), &payload); err != nil {
			fmt.Println("Error decoding JSON:", err)
			responseOnError(conn)
		}

		fmt.Println("Decoded payload:", payload)

		// Validate the method field
		if payload.Method != "isPrime" {
			fmt.Println("Invalid method:", payload.Method)
			responseOnError(conn)
		}

		// Let's leave the primality testing for another day :)
		response := Response{Method: "isPrime", Prime: primes.IsPrime(payload.Number)}

		// Marshal the response to JSON
		responseData, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error encoding JSON response:", err)
			return
		}

		// Send a message back
		_, err = conn.Write(append(responseData, '\n'))
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from connection:", err)
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
