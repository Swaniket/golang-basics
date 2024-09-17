package main

import "fmt"

var jwtToken string = "sasasasa"

const LoginToken string = "Some token here" // using 'L' makes it a public variable

func main() {
	// var username string = "Swaniket"
	// fmt.Println(username)
	// fmt.Printf("Variable is of type: %T \n ", username)

	// var isLoggedIn bool = false
	// fmt.Println(isLoggedIn)
	// fmt.Printf("Variable is of type: %T \n ", isLoggedIn)

	// var smallVal uint8 = 255
	// fmt.Println(smallVal)
	// fmt.Printf("Variable is of type: %T \n ", smallVal)

	// var smallVal int = 25523232
	// fmt.Println(smallVal)
	// fmt.Printf("Variable is of type: %T \n ", smallVal)

	// var smallFloat float32 = 256.877197876671
	// fmt.Println(smallFloat)
	// fmt.Printf("Variable is of type: %T \n ", smallFloat)

	var largeFloat float64 = 256.877197876671 // float64 is more presise
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n ", largeFloat)

	// Default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n ", anotherVariable)

	// Implicit Type
	var website = "www.google.com"
	fmt.Println(website)

	// We can also declare variables w/o using the var keyword
	// := is known as Walrus Operator (Only allowed inside of any method)
	noOfUsers := 10.09
	fmt.Println(noOfUsers, LoginToken, jwtToken)
}
