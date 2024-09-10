package main

import "fmt"
import "os"
import "bufio"


func check(e error){
	if e != nil{
		panic(e)
	}
}


func main(){

	// Get file name
	fmt.Printf("Input filemane: ")
	inputFileName := ""
	fmt.Scanf("%s", &inputFileName)
	
	// Open input file
	inputFile, err_in := os.Open(inputFileName)
	check(err_in)
	defer inputFile.Close()
	
	
	// Scan input file and write output file
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

}
