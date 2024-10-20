package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platfrom string   `json:"website"`
	Password string   `json:"-"` // giving a -  means it wont come when someone consumes the API
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")

	// Enncoding of JSON (I have a data, I want to convert the data into valid JSON)
	// EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {
	lcoCourses := []course{
		{"Go Bootcamp", 299, "youtube.com", "abc123", []string{"Web Dev", "Backend"}},
		{"ReactJS", 399, "youtube.com", "bcd123", []string{"Web Dev", "Frontend"}},
		{"MEAN", 199, "youtube.com", "jhy123", nil},
	}

	// Package the data as JSON data
	// finalJSON, err := json.Marshal(lcoCourses)

	// To convert JSON with indentation
	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJSON)
}

func DecodeJSON() {
	mockJSON := []byte(`
        {
                "courseName": "Go Bootcamp",
                "Price": 299,
                "website": "youtube.com",
                "tags": [
                        "Web Dev",
                        "Backend"
                ]
        }
`)

	var lcoCourse course

	// Checking if the JSON data is valid or not
	checkValid := json.Valid(mockJSON)

	if checkValid {
		fmt.Println("The JSON is Valid")
		json.Unmarshal(mockJSON, &lcoCourse) // Passing the referance of the interface
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("The JSON is Invalid!")
	}

	// Add data in key-value pair
	var myOnlineData map[string]interface{} // We are not sure what value will come up, so we can create an interface and auto populate it
	json.Unmarshal(mockJSON, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	// Looping through the data
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T \n", k, v, v)
	}
}
