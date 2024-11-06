package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course (Should go to a different file)
type TCourse struct {
	CourseId    string   `json:"courseId"`
	CourseName  string   `json:"courseName"`
	CoursePrice int      `json:"price"`
	Author      *TAuthor `json:"author"` // As the TAuthor struct is not created yet, we are passing a pointer to this custom struct
}

type TAuthor struct {
	Fullname string `json:"fullName"`
	Website  string `json:"website"`
}

// Fake DB
var courses []TCourse

// Helper Methods (Should go to a different file)
func (course *TCourse) IsEmpty() bool {
	// return course.CourseId == "" && course.CourseName == "" // return true if both CourseId and CourseName is empty
	return course.CourseName == "" // Not checking for courseId becuase we want to manually create the courseID
}

// Controllers (Should go to a different file)

// Serve Home Route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the API home route</h1>"))
}

func getAllCourses(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Inside get all courses")

	// Set the headers
	writer.Header().Set("Content-Type", "application/json")
	// Return all the data in the fake DB as JSON
	json.NewEncoder(writer).Encode(courses)
}

func getCourseByCourseId(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// Grab the courseId from request
	params := mux.Vars(request)

	// Step 1 - Loop though the courses
	for _, course := range courses {
		// Step 2 - Find the matching ID and return the response
		if course.CourseId == params["id"] {
			json.NewEncoder(writer).Encode(course)
			return
		}
	}

	json.NewEncoder(writer).Encode("No course found with the given ID")
}

func createCourse(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// Case 1: Body is empty
	if request.Body == nil {
		json.NewEncoder(writer).Encode("Invalid request body")
	}

	// Case 2: Body is empty

	var newCourse TCourse
	// Decoding the value in the request.body
	_ = json.NewDecoder(request.Body).Decode(&newCourse)

	if newCourse.IsEmpty() {
		json.NewEncoder(writer).Encode("The JSON is empty")
		return
	}

	// Generate a unique ID & convert them in string
	rand.Seed(time.Now().UnixNano())
	newCourse.CourseId = strconv.Itoa(rand.Intn(100))

	// Add the couse in the slice
	courses = append(courses, newCourse)

	// Send the data
	json.NewEncoder(writer).Encode(newCourse)
	return
}

func updateCourseById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// Step 1 - Grab Id from request
	params := mux.Vars(request)

	// Loop through the courses, get the course by ID, remove it and add with the give ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// Remove
			courses = append(courses[:index], courses[index+1:]...)

			// Create the course Ref
			var updatedCourse TCourse

			// Decode the request body and add the ID
			_ = json.NewDecoder(request.Body).Decode(&updatedCourse)
			course.CourseId = params["id"]

			// Append the updated course to the courses slice
			courses = append(courses, updatedCourse)

			// Sending back response
			json.NewEncoder(writer).Encode(updatedCourse)
			return
		}

		json.NewEncoder(writer).Encode("The record with the specified ID does not exists")
	}
}

func deleteOneCourse(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// Step 1 - Grab Id from request
	params := mux.Vars(request)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(writer).Encode("Delete Successful")
			break
		}
	}
}

func main() {
	fmt.Println("Working with APIs")

	// Creating the router
	router := mux.NewRouter()

	// Seeding initial data
	courses = append(courses, TCourse{CourseId: "432", CourseName: "Go with Golang", CoursePrice: 299, Author: &TAuthor{Fullname: "Swaniket Chowdhury", Website: "swaniket.com"}})
	courses = append(courses, TCourse{CourseId: "209", CourseName: "ReactJS", CoursePrice: 199, Author: &TAuthor{Fullname: "Hitesh Chowdhury", Website: "learn.lco.dev"}})

	// Routing
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getCourseByCourseId).Methods("GET")
	router.HandleFunc("/course", createCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateCourseById).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":4000", router))

}
