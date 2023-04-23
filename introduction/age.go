package main

import "fmt"

func main() {
	var age int

	fmt.Println("Enter your age:")
	fmt.Scan(&age)

	month := age * 12
	days := age * 365

	fmt.Println("Your age is", age)
	fmt.Println("Your age in months is", month)
	fmt.Println("Your age in days is", days)
}

// This program is a simple age calculator.
// The possible development of this program is to add funcionality that will check leap years and add the days of the leap years to the total days.
// Another one is to check the current date and add the days of the current year to the total days.