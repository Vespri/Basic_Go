package main

import "fmt"

func main() {
	var name = "Kresna"
	var midName = "Vespri"
	var lastName = "Wijaya"
	var newlength = len(midName)

	switch name {
	case "Kresna":
		fmt.Println("Hi", name)
	case "Vespri":
		fmt.Println("Hi", name)
	default:
		fmt.Println("Who is there ?")
	}

	// Normal/Common Statement
	var length = len(name)
	switch length > 10 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama diterima !")
	default:
		fmt.Println("Masukkan tidak sesuai")
	}

	// Short Statement
	switch newLength := len(lastName); newLength > 10 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama diterima !")
	default:
		fmt.Println("Masukkan tidak sesuai")
	}

	// Switch without Expression
	switch {
	case newlength > 10:
		fmt.Println("Nama terlalu panjang")
	case newlength > 7:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama diterima !")
	}
}
