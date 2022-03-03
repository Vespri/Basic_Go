package main

import "fmt"

func main() {
	// Single constant declaration
	const firstName string = "Kresna Vespri"
	const lastName = "Wijaya"

	fmt.Println("First Name :", firstName)
	fmt.Println("Last Name :", lastName)

	// Multiple constant declaration
	const (
		Nickname = "Kresna"
		fullName = "Kresna Vespri Wijaya"
	)

	fmt.Println("Nickname :", Nickname)
	fmt.Println("Full Name :", fullName)
}
