package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const POSTS_URL = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("Making API Calls")

	GETCall(POSTS_URL)
}

func GETCall(url string) {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// Close the requests
	defer response.Body.Close()

	// Handling the response
	fmt.Println("Stats Code: ", response.StatusCode)
	responseContent, _ := io.ReadAll(response.Body)
	// fmt.Println("responseContent: ", string(responseContent))

	// We can also transform the string using the strings package (Preffered)
	var responseString strings.Builder

	byteCount, _ := responseString.Write(responseContent) // This will give me the count
	fmt.Println("Byte Count: ", string(byteCount))

	stringResponse := responseString.String()
	fmt.Println("stringResponse: ", stringResponse)
}

func POSTCall(url string) {

}
