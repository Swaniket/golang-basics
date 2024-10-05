package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops!")

	days := []string{"Sunday", "Monday", "Friday", "Saturday"}

	// Go only have for loops - there are different variations of for loops

	// Traditional loop
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// Using the range - range auromatically returns the index
	for i := range days {
		fmt.Println("i", i)
		fmt.Println(days[i])
		// fmt.Println(d)
	}

	// Range function also return the index & value
	for index, day := range days {
		fmt.Println("index", index)
		fmt.Println("day", day)
	}

	// We can also mimic the behavior of while loop
	randomValue := 1

	for randomValue < 10 {

		if randomValue == 2 {
			goto swaniket
		}

		if randomValue == 5 {
			randomValue++
			continue
		}
		// if randomValue == 5 {
		// 	break
		// }

		fmt.Println("Value is:", randomValue)
		randomValue++
	}

	// We can also use goto
swaniket:
	fmt.Println("My name is swaniket")

}
