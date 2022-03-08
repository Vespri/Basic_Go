package main

import "fmt"

func main() {
	var days = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Friday",
		"Saturday",
		"Sunday",
	}
	var slice1 = days[1:4]
	fmt.Println(slice1)

	fmt.Print("\n")

	// Function slice
	// len (length)
	fmt.Println(len(slice1))
	// cap (capacity)
	fmt.Println(cap(slice1))

	fmt.Print("\n")

	// append (new slice)
	var slice2 = days[5:]
	fmt.Println(slice2) // [Saturday Sunday]

	var slice3 = append(slice2, "Holiday")
	fmt.Println(slice3) // [Saturday Sunday Holiday]
	slice3[0] = "Kresna"
	fmt.Println(slice3) // [Kresna Sunday Holiday]

	fmt.Println(slice2) // [Saturday Sunday]
	fmt.Println(days)   // [Senin Selasa Rabu Kamis Friday Saturday Sunday]

	fmt.Print("\n")

	var slice4 = days[3:6]
	fmt.Println(slice4) // [Kamis Friday Saturday]

	var slice5 = append(slice4, "Busy")
	fmt.Println(slice5) // [Kamis Friday Saturday Busy]
	slice5[1] = "Kresna"
	fmt.Println(slice5) // [Kamis Kresna Saturday Busy]

	fmt.Println(slice4) // [Kamis Kresna Saturday]
	fmt.Println(days)   // [Senin Selasa Rabu Kamis Kresna Saturday Busy]

	fmt.Print("\n")

	// make slice
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Kresna"
	newSlice[1] = "Vespri"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Print("\n")

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	fmt.Print("\n")

	// Array vs Slice
	Array := [...]int{1, 2, 3, 4}
	Slice := []int{1, 2, 3, 4}

	fmt.Println("This is an array :", Array)
	fmt.Println("This is a slice :", Slice)
}
