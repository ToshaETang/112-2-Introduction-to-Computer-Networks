package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:11993")
	if err != nil {
		fmt.Println("Unable to connect to the server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to the server")

	// Send data
	message := "76\n" // "50\n"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	fmt.Println("Data sent:", message)

	// Receive and print server response
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("Server response:", scanner.Text())
		break // Print only one line of response
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading server response:", err)
		return
	}

	// Close the connection
	err = conn.Close()
	if err != nil {
		fmt.Println("Error closing connection:", err)
		return
	}

	fmt.Println("Connection closed")
}

