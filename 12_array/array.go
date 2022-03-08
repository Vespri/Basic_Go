package main

import "fmt"

func main() {
	// Array 1
	var names [3]string

	names[0] = "Kresna"
	names[1] = "Vespri"
	names[2] = "Wijaya"

	fmt.Println("Data pertama :", names[0])
	fmt.Println("Data kedua :", names[1])
	fmt.Println("Data ketiga :", names[2])
	fmt.Println("All data :", names)

	// Array 2
	var values = [4]int{
		12,
		14,
		16,
		18,
	}
	fmt.Println("Data pertama :", values[0])
	fmt.Println("Data kedua :", values[1])
	fmt.Println("Data ketiga :", values[2])
	fmt.Println("Data ketiga :", values[3])
	fmt.Println("All data :", values)
}
