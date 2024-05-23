package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

func CuustomNotFund(fs http.FileSystem) http.Handler{
	fsh := http.FileServer(fs);
	
	return http.HandlerFunc(  func(w http.ResponseWriter, r *http.Request){
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err){
			fmt.Fprintf(w,"File Not Found\n")
			return
		}
		fsh.ServeHTTP(w, r)
	}  )

}

func main(){
 	
	fmt.Println("Launching server...")
	
	////http.ListenAndServe(":12017", CuustomNotFund(http.Dir(".")))
	// 啟動 TLS 伺服器，監聽指定的端口，使用指定的證書和私鑰文件
	http.ListenAndServeTLS(":12017", "server.cer", "server.key", CuustomNotFund(http.Dir(".")))

}

//openssl genrsa -out server.key 2048
//openssl req -new -x509 -key server.key -out server.cer -days 365

//curl -k https://127.0.0.1:12017/qwerty.htm



