package main

import "fmt"

func main() {
	var name1 = "Kresna"
	var name2 = "Kresna"
	var name3 = "Vespri"

	var result1 bool = name1 == name2
	var result2 bool = name1 == name3

	fmt.Println(name1, "=", name2, ":", result1)
	fmt.Println(name1, "=", name3, ":", result2)

	var value1 = 10
	var value2 = 20

	// ==
	fmt.Println(value1, "==", value2, ":", value1 == value2)
	// !=
	fmt.Println(value1, "!=", value2, ":", value1 != value2)
	// >
	fmt.Println(value1, ">", value2, ":", value1 > value2)
	// <
	fmt.Println(value1, "<", value2, ":", value1 < value2)
}
