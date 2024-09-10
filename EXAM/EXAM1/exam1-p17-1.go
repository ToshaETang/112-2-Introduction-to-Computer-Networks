package main

import "fmt"
import "os"


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
	
	fi, err := os.Stat(inputFileName)
	if err == nil{
		fmt.Println("file size : ",fi.Size())
	}
	
}
	
