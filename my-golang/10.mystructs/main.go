package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")

	// There are no classes in Go, instead we use structs
	// There are no inheritane and super and parent concept in Go

	swaniket := User{"Swaniket", "swaniket@google.com", true, 27}
	fmt.Println("swaniket", swaniket)
	fmt.Printf("Swaniket's details are %+v \n", swaniket) // %+v is used to print out any structures

	// We can als access single properties
	fmt.Printf("Email Id is: %v", swaniket.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
