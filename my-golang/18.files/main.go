package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files")

	content := "This needs to go in a file"

	// Create Files
	file, err := os.Create("./testFile.txt")
	checkError(err)

	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Println("Length is", length)
	defer file.Close()

	readFile("./testFile.txt")
}

// Reading a file
func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName) // It doesn't read in a string format, its in byte format
	checkError(err)

	fmt.Println("Raw Data is \n", databyte)
	fmt.Println("Text Data is \n", string(databyte))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
