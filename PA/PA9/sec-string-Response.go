package main

import "fmt"
import "bufio"
import "net/http"
import "crypto/tls"

// 用於檢查錯誤的函數，如果有錯誤發生則會引發panic
func check(e error) {
 if e != nil {
 panic(e)
 }
}

func main() {
 // 載入伺服器的證書和私鑰
 cert, _ := tls.LoadX509KeyPair("server.cer", "server.key")
 
 // 配置TLS的設置，包括證書
 config := tls.Config{Certificates: []tls.Certificate{cert}}
 
 fmt.Println("Launching server...")
 
 // 開始在指定的端口上監聽TLS連接
 ln, _ := tls.Listen("tcp", ":<your port#>", &config)
 defer ln.Close()
 
 // 接受來自客戶端的連線
 conn, _ := ln.Accept()
 defer conn.Close()
 
 // 使用bufio讀取客戶端發送的請求
 reader := bufio.NewReader(conn)
 req, _ := http.ReadRequest(reader)
 
 // 輸出請求的方法（例如GET, POST等）
 fmt.Printf("Method: %s\n", req.Method)
 
 // 向客戶端回應404 Not Found錯誤
 fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n")
 fmt.Fprintf(conn, "Date: ...\r\n")
 fmt.Fprintf(conn, "\r\n")
 fmt.Fprintf(conn, "File not found\r\n")
 fmt.Fprintf(conn, "\r\n") 
}
