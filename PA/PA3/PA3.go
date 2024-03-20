package main
import "fmt"
import "bufio"
import "net"
import "os"
import "io"
import "strconv"

func check(e error){
	if e != nil {
		panic(e)
	}
}

 
func main(){
	
	// connect to port 12000
	conn, errc := net.Dial("tcp", "127.0.0.1:12000")
 	check(errc)
 	defer conn.Close()


	// Get file name
	fmt.Printf("Input filemane: ")
	inputFileName := ""
	fmt.Scanf("%s", &inputFileName)
	
	
	// Open input file
	inputFile, err_in := os.Open(inputFileName)
	check(err_in)
	defer inputFile.Close()
	
	
	// read and send file size
	writer := bufio.NewWriter(conn)
	fi, err := os.Stat(inputFileName)
	
	if err == nil{ // file exist
		fmt.Println("Send the file size first: ",fi.Size())
		fs := strconv.FormatInt(fi.Size(), 10) // int64 -> string
		
		writer.WriteString(fs) // send file size
		writer.WriteString("\n")
		writer.Flush()
 		
		io.Copy(conn, inputFile) // send file content
		
		// get message from server
		scanner := bufio.NewScanner(conn)
		if scanner.Scan() {
 			fmt.Println("Server says: ",scanner.Text()) // message from server
 		}
	}else{
		fmt.Printf("No File!!!!!!!")
	}

}
