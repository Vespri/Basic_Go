package main

import "fmt"

func main() {
	// Array 1
	var names [3]string

	names[0] = "Kresna"
	names[1] = "Vespri"
	names[2] = "Wijaya"

	fmt.Println("First data :", names[0])
	fmt.Println("Second data :", names[1])
	fmt.Println("Third data :", names[2])
	fmt.Println("All data :", names)

	// Array 2
	var values = [4]int{
		12,
		14,
		16,
		18,
	}
	fmt.Println("First data :", values[0])
	fmt.Println("Second data :", values[1])
	fmt.Println("Third data :", values[2])
	fmt.Println("Fourth data :", values[3])
	fmt.Println("All data :", values)

	// Function Array
	// len
	fmt.Println("Length names's array = ", len(names))
	fmt.Println("Length values's array = ", len(values))
	// change array data
	values[0] = 11
	fmt.Println("All data :", values)

}
