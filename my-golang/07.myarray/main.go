package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")
	// Note: Arrays are not much used in golang because its underpowered

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Mango"

	fmt.Println("Fruits", fruitList)
	fmt.Println("Fruits", len(fruitList)) // The output will be 4 even though I have 3 elements
	// The reason being go lang automatically inserts a '' at the missing spot

	var numList [4]int

	numList[0] = 100
	numList[3] = 400

	fmt.Println("Nums", numList) // For number it inserts 0
	fmt.Println("Nums", len(numList))

	// Init with values
	var vegList = [3]string{"Potato", "Beans", "Mashroom"}
	fmt.Println("vegList", vegList)
	fmt.Println("vegList", len(vegList))
}
