// When functions go into classes / structs we call them methods

package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// // oneAge will not be exportable because the first letter is not capital
	// oneAge int
}

// We need to pass a struct or a portery of a struct to a method
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

func main() {
	fmt.Println("Welcome to methods")

	swaniket := User{"Swaniket", "swaniket@google.com", true, 27}
	fmt.Println("swaniket", swaniket)
	fmt.Printf("Swaniket's details are %+v \n", swaniket) // %+v is used to print out any structures

	// We can als access single properties
	fmt.Printf("Email Id is: %v \n", swaniket.Email)

	swaniket.GetStatus()
	swaniket.NewMail()

	fmt.Printf("Email Id is: %v \n", swaniket.Email) // This will print swaniket@google.com because NewMail() func manupulates a copy of the struct
}
