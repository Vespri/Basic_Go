package main

import "fmt"

func getName(name string) string {
	if name == "" {
		return "Hello world !"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getName("Kresna")
	fmt.Println(result)

	fmt.Println(getName(""))
}
