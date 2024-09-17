package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Tomato", "Mango"} // If we use this syntax then we need to init
	fmt.Printf("Type of fruit list %T\n", fruitList)
	fmt.Println("fruitList", fruitList)

	// Appending values - into new variable (We can assign the value into the same array as well)
	updatedFruitList := append(fruitList, "Banana", "Peach")
	fmt.Println("updatedFruitList", updatedFruitList)

	// Slicing the slices
	// updatedFruitList = append(updatedFruitList[1:]) // [Tomato Mango Banana Peach]
	// updatedFruitList = append(updatedFruitList[1:2]) // [Tomato]

	// We don't need to use append to slice
	updatedFruitList = updatedFruitList[1:2] // [Tomato]
	fmt.Print("updatedFruitList sliced", updatedFruitList)

	// Using make() to create a slice
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 // The array was of size 4 and we added a 5th element - so the program will crash with a index out of range error

	// But if we use the append method, the memory is reallocated and we can store the values
	highScores = append(highScores, 555, 666, 321)

	fmt.Println("highScores", highScores)
	fmt.Println("Is Sorted", sort.IntsAreSorted(highScores))
	sort.Ints(highScores) // Sort an integer slice
	fmt.Println("Sorted highScores", highScores)
	fmt.Println("Is Sorted", sort.IntsAreSorted(highScores))

	// Remove value from slice based of index
	fmt.Println("-------------------------------------------------------------------------------")
	var courses = []string{"ReactJS", "Javascript", "Typescript", "Swift", "Python", "Ruby", "Go"}
	fmt.Println("courses", courses)

	var indexToRemove int = 2

	// We can also remove element using .append()
	courses = append(courses[:indexToRemove], courses[indexToRemove+1:]...) // will learn ... later
	fmt.Println("courses after remove", courses)
}
