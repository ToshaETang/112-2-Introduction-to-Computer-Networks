package main
import (
    "fmt"
    "strconv"
    "net"
    "bufio"
    "os"
    "time"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}


func handleConnection(conn net.Conn){
	reader := bufio.NewReader(conn)
	message, errr := reader.ReadString('\n')
	check(errr)

	//--------------------------------------------------
	//【get content from client】
	
	size := 0;
	tmp := message[:len(message)-1]
	inputSize, err := strconv.Atoi(tmp)
	fmt.Printf("upload file size: %d\n", inputSize)
		
		
	// Open whatever.txt file
	whatever, err_in := os.Create("whatever.txt")
	check(err_in)
	defer whatever.Close()
	writer := bufio.NewWriter(whatever)
	
	
	check(err)
	for i := 1; i < 1000; i++ {
		message2, errr := reader.ReadString('\n')
  	check(errr)
    size = size + len(message2)
    	
    tmp2 := fmt.Sprintf("%d %s", i, message2)
    writer.WriteString(tmp2) // save sentence to writer	
    writer.Flush() // write in file
		
    if size == inputSize {
    	break
    }
	}
	//writer.Flush() // write in file
	
	fileinfo, err := whatever.Stat()
  check(err)
  outputSize := fileinfo.Size()
	fmt.Printf("outload file size: %d\n", outputSize)

	
	//---------------------------------------------------
	//【send message to client】
 	writer2 := bufio.NewWriter(conn)
 	newline := fmt.Sprintf("%d bytes received, %d bytes file generated\n",inputSize, outputSize)
 	_, errw := writer2.WriteString(newline)
 	check(errw)
 	writer2.Flush()
 	time.Sleep(5 * time.Second)

}




func main() {
	
	//--------------------------------------------------
	//【get size from client】
	
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12017")
	defer ln.Close()
	
	for{
	
		conn, _ := ln.Accept()
		defer conn.Close()
		go handleConnection(conn)
	
 	}//for loop
 	
 	
}

