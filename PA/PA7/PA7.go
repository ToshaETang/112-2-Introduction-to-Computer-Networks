package main
import ("fmt"
				"bufio"
				"net"
				"net/http"
				"strings"
				"os"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {

	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12017")
	defer ln.Close()
	
	for{
	
		conn, _ := ln.Accept()
	
		reader := bufio.NewReader(conn)
	
		req, err := http.ReadRequest(reader)
		check(err)
		
		//get file name in URL
		token := strings.Split(req.RequestURI, "/")
		fileName := token[1]
		//fmt.Printf("RequestURI: %s\n", req.RequestURI)
		//fmt.Printf("fileName: %s\n", fileName)
		
		// check if the file exist
		fileInfo, err := os.Stat(fileName)
		if err != nil {
			fmt.Println("File not found")

		} else{
			fmt.Println("File size = ", fileInfo.Size())
		}
		
		conn.Close()

	}//for
	
	
}//main




