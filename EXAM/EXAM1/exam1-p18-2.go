package main

import (
	"fmt"
	"net"
)

func main() {
	
	conn, err := net.Dial("tcp", "localhost:11992")
	if err != nil {
		fmt.Println("NO", err)
		return
	}
	defer conn.Close()

	fmt.Println("CCONNECt")

	message := "50\n"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("WRONG", err)
		return
	}

	fmt.Println("SEND", message)

	err = conn.Close()
	if err != nil {
		fmt.Println("WRONG", err)
		return
	}

	fmt.Println("STOP")
}

