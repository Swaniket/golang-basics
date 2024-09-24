package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["TS"] = "Typescript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println("Languages", languages) // OP - map[GO:Golang JS:Javascript PY:Python TS:Typescript]
	// NOTE: These are not comma separated values

	fmt.Println("Grabbing one value JS:", languages["JS"])

	// Deleting some value
	delete(languages, "PY")
	fmt.Println("Languages", languages)

	// Looping though the maps
	for key, value := range languages {
		// fmt.Println("KEY", key)
		// fmt.Println("VALUE", value)
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}
