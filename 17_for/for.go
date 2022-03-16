package main

import "fmt"

func main() {
	// For Loops
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	fmt.Print("\n")

	// For Loops with Statement
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke", i)
	}

	fmt.Print("\n")

	// For Slice
	slice := []string{"Kresna", "Vespri", "Wijaya"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Print("\n")

	// For Range(Slice)
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
		// fmt.Println(value)
	}

	fmt.Print("\n")

	// For Range(Map)
	names := make(map[string]string)
	names["firstName"] = "Kresna"
	names["midName"] = "Vespri"
	names["lastName"] = "Wijaya"

	for key, value := range names {
		fmt.Println(key, "=", value)
	}
}
