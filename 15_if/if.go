package main

import "fmt"

func main() {
	var name = "Kresna"
	var fullName = "Kresna Vespri Wijaya"

	if name == "Vespri" {
		fmt.Println("Hi", name)
	} else if name == "Kresna" {
		fmt.Println("Hi", name)
	} else {
		fmt.Println("Who is there ?")
	}

	// Normal/Common Statement
	var length = len(name)

	if length > 10 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	// Short Statement
	if newLength := len(fullName); newLength > 30 {
		fmt.Println("Nama lengkap terlalu panjang")
	} else {
		fmt.Println("Nama lengkap diterima !")
	}
}
