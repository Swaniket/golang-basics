package main

import "fmt"

func main() {
	fmt.Println("Welcome to if-else")

	loginCount := 9
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount >= 10 && loginCount < 20 {
		result = "Super User"
	} else {
		result = "Super Active User"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// We can also initialize something and apply if-else to it
	// This is useful when for example the value comes from a web request
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}
}
