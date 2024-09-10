package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	
	outputFileName := "whatever.txt"
	// Check file name
	//fmt.Printf("%s %s\n", inputFileName, outputFileName)
	
	
	// Open input file
	inputFile, err_in := os.Open(inputFileName)
	check(err_in)
	defer inputFile.Close()
	
	
	// Create output file
	outputFile, err_out := os.Create(outputFileName)
	check(err_out)
	defer outputFile.Close()
	
	
	// Scan input file and write output file
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	for scanner.Scan(){
		line := scanner.Text()
		tmp := fmt.Sprintf("%s\n", strings.ToUpper(line)) //int -> str, then merge
		writer.WriteString(tmp)
		writer.Flush()
	}

}
