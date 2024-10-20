package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const POSTS_URL = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("Making API Calls")

	// GETCall(POSTS_URL)
	// POSTCall(POSTS_URL)
	PerformPostForm(POSTS_URL)
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
	// Fake JSON payload
	reqBody := strings.NewReader(`
		{
			"userId": 9999,
			"id": 999,
			"title": "This is a test post",
			"body": "This is the body of the test post"
		}
	`)

	response, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseContent, _ := io.ReadAll(response.Body)
	fmt.Println("responseContent: ", string(responseContent))
}

func PerformPostForm(reqUrl string) {
	// Building Form Data
	data := url.Values{}
	data.Add("firstName", "Swaniket")
	data.Add("lastName", "Chowdhury")
	data.Add("email", "swaniket@go.dev")

	// It creates a special post request which is URL encoded
	response, err := http.PostForm(reqUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseContent, _ := io.ReadAll(response.Body)
	fmt.Println("responseContent: ", string(responseContent))
}
