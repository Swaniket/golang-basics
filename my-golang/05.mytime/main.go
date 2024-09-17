package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("Welcome to time program")

	presentTime := time.Now()

	fmt.Println("Present time: ", presentTime)
	// fmt.Println("Present time + 1", presentTime.After(1))

	fmt.Println("-------------- Formatted Time --------------")
	fmt.Println(presentTime.Format("01-02-2006")) // We need to give 01-02-2006 to format in month-date-year
	fmt.Println(presentTime.Format("02/01/2006")) // Interchanging the month with date and using a different separator
	// We can also format the date additionally by passing day and time
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	// We need to give the month from time package becase the type for month is time.month
	// Function Signature: time.Date(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location)
	createdDate := time.Date(2024, time.September, 29, 02, 01, 00, 00, time.Local)
	fmt.Println("Created Time", createdDate)

}
