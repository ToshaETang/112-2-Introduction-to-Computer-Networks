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
	
	fmt.Printf("Output filemane: ")
	outputFileName := ""
	fmt.Scanf("%s", &outputFileName)
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
	count := 1 // use it to count line
	for scanner.Scan(){
		tmp := fmt.Sprintf("%d %s\n", count, scanner.Text()) //int -> str, then merge
		writer.WriteString(tmp)
		writer.Flush()
		count++
	}

}
