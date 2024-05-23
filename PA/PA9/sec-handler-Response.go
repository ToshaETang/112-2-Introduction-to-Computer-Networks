package main

import (
	"fmt"
	"net/http"
)

// helloHandler 處理 /hello 路徑的請求，回應 "Hello, world!"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func main() {
	// 輸出伺服器啟動信息
	fmt.Println("Launching server...")

	// 將 helloHandler 包裝成 http.Handler
	hh := http.HandlerFunc(helloHandler)
	// 將 /hello 路徑的請求交給 helloHandler 處理
	http.Handle("/hello", hh)

	// 使用 FileServer 提供靜態文件服務，根目錄為當前目錄
	fs := http.FileServer(http.Dir("."))
	// 將根路徑的請求交給 FileServer 處理，並去掉路徑中的前綴
	http.Handle("/", http.StripPrefix("/", fs))

	// 啟動 TLS 伺服器，監聽指定的端口，使用指定的證書和私鑰文件
	http.ListenAndServeTLS(":<your port#>", "server.cer", "server.key", nil)
}

