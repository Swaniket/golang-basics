package main

import "fmt"

func main() {
	// Pointers gives us the gareetee that the operation is performed in the actual value
	fmt.Println("Pointers")

	// Normally we define a value like
	// var num int = 2
	// fmt.Println(num)

	// We can create a pointer that points to an integer
	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)

	// Using a pointer
	myNumber := 23
	var newPtr = &myNumber // We are creating a pointer which is also referancing to something. Thats why we are using &

	fmt.Println("Value of the pointer ref", newPtr) // This will return the memory address something like 0xc00008c0d0
	fmt.Println("Value of actual pointer", *newPtr) // We add * if we want to access the value

	// We can also do something like this
	*newPtr = *newPtr * 2
	fmt.Println("Updated value is", myNumber)
}
