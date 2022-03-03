package main

import "fmt"

func main() {
	// Variable declaration with var & data type
	var name string

	name = "Kresna"
	fmt.Println("Nickname :", name)

	// Variable declaration with var & without data type (String)
	var fullName = "Kresna Vespri Wijaya"
	fmt.Println("Full Name :", fullName)

	// Variable declaration with var & without data type (Integer)
	var age = 21
	fmt.Println("Age :", age)

	// Variable declaration without var & data type (First time)
	country := "Indonesia"
	fmt.Println("Country :", country)

	// Variable declaration without var & data type (Second time & next)
	country = "Singapore"
	fmt.Println("Country :", country)

	// Multiple variable declaration
	var (
		firstName = "Kresna Vespri"
		lastName  = "Wijaya"
	)
	fmt.Println("First Name :", firstName)
	fmt.Println("Last Name :", lastName)
}
