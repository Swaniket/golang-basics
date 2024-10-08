package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	response, err := http.Get(URL)
	handleError(err)

	fmt.Printf("Response Type %T\n", response)

	// We need to close every request we make
	defer response.Body.Close()

	// We need to use io.ReadAll() to read the response
	dataByte, err := io.ReadAll(response.Body)
	handleError(err)

	fmt.Println(string(dataByte))
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
