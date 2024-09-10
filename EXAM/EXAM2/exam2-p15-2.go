package main
import "fmt"
import "bufio"
import "net"
func check(e error) {
 if e != nil {
 panic(e)
 }
}
func main() {
 fmt.Println("Launching server...")
 ln, _ := net.Listen("tcp", ":12000")
 conn, _ := ln.Accept()
 defer ln.Close()
 defer conn.Close()
 
 scanner := bufio.NewScanner(conn)
 message := ""
 if scanner.Scan() {
 message = scanner.Text()
 fmt.Println(message)
 } 

}


