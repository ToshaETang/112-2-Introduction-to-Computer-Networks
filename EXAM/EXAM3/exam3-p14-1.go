
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
	http.ListenAndServe(":31701", CuustomNotFund(http.Dir(".")))
}

