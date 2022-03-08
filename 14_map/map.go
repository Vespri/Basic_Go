package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Kresna",
		"address": "Palembang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])
	fmt.Println(len(person))

	fmt.Print("\n")

	// make map
	var repo = make(map[string]string)
	repo["title"] = "Basic-Go"
	repo["owner"] = "Vespri"
	repo["website"] = "Github"

	// before delete
	fmt.Println(repo)
	fmt.Println(len(repo))

	// after delete
	delete(repo, "website")
	fmt.Println(repo)
	fmt.Println(len(repo))
}
