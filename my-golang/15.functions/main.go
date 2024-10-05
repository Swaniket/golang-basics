package main

import "fmt"

func main() {
	// We can't have another function inside of a function
	fmt.Println("Welcome to functions")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	result2, str := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Result2 is: ", result2, str)
}

// We need to mention the types for the arguments & also for the function if its returning anything
func adder(num1 int, num2 int) int {
	return num1 + num2
}

// Pro function - We can pass multiple number of args
// The args will be stored as a slice
// ... is known as variatic
func proAdder(values ...int) (int, string) {
	total := 0

	for _, v := range values {
		total += v
	}

	// We can also return multiple things
	return total, "Hello Testing"
}

func greeter() {
	fmt.Println("Namaste from Golang")
}
