package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our app")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks your your rating", input)

	// Convert into number
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // 64 is the byte size

	// Handling the error
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	} else {
		newRating := numRating + 1
		fmt.Println("Added 1 to your rating", newRating)
	}
}
