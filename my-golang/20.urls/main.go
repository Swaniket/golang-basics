package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=232ewew2232eewg"

func main() {
	fmt.Println("Welcome to URLs")
	fmt.Println("myUrl", myUrl)

	// Parsing the URL
	result, _ := url.Parse(myUrl)

	fmt.Println("Protocol", result.Scheme)
	fmt.Println("Host", result.Host)
	fmt.Println("Path", result.Path)
	fmt.Println("Port", result.Port())
	fmt.Println("Query", result.RawQuery)

	// Extracting the query params in a better way
	qParams := result.Query()
	fmt.Printf("Type of QParams %T\n", qParams)
	fmt.Println("QParams", qParams)
	fmt.Println("Course name", qParams["coursename"])
	fmt.Println("Payment Id", qParams["paymentid"])

	// Loop over it
	for _, val := range qParams {
		fmt.Println("Params are: ", val)
	}

	// Construct a URL
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=swaniket",
	} // IMPORTANT - We always need to pass on a reference of it, not a copy

	fmt.Println("New URL", partsOfUrl.String())
}
