

package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

func CustomNotFound(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		if r.URL.Path == "/hello" {
			fmt.Fprintf(w, "Hi~\n")
			return
		}

		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			fmt.Fprintf(w, "Hmm.. not here.\n")
			return
		}
		fsh.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Launching server...")

	
	http.ListenAndServeTLS(":31701", "server.cer", "server.key", CustomNotFound(http.Dir(".")))
}


