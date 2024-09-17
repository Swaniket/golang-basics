package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMsg := "Welcome to user inpur"

	fmt.Println(welcomeMsg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza")

	// Comma ok syntax or Error ok syntax (Replacement for Try Catch)
	input, err := reader.ReadString('\n')
	fmt.Println("Your rating is,", input)
	fmt.Printf("Type of the rating is %T", input)

	fmt.Println("The Error is", err)

}
