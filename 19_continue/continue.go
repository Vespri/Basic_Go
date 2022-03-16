package main

import "fmt"

func main() {
	// For Continue(even)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}

	fmt.Print("\n")

	//  For Continue(odd)
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}
