package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:11991")
	if err != nil {
		fmt.Println("NO", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connecct !!!")

	err = conn.Close()
	if err != nil {
		fmt.Println("wrong", err)
		return
	}

	fmt.Println("Stop !!!")
}

