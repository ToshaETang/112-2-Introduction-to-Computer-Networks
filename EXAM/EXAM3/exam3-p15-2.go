package main

import (
    "bufio"
    "crypto/tls"
    "fmt"
    "log"
)

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func main() {
 
    cert, err := tls.LoadX509KeyPair("client.cer", "client.key")
    check(err)

  
    config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

    
    conn, err := tls.Dial("tcp", "localhost:12000", &config)
    check(err)
    defer conn.Close()

    
    _, err = conn.Write([]byte("PLAY\n"))
    check(err)

   
    reader := bufio.NewReader(conn)
    response, err := reader.ReadString('\n')
    check(err)


    fmt.Print("Server response: ", response)
}

