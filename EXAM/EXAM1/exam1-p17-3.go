package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Get file name
	fmt.Printf("Input filename: ")
	var inputFileName string
	fmt.Scanln(&inputFileName)

	// Open input file
	inputFile, err := os.Open(inputFileName)
	check(err)
	defer inputFile.Close()

	// Scan input file and print uppercase text
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		// Convert the line to uppercase and print it
		fmt.Println(strings.ToUpper(line))
	}
}

