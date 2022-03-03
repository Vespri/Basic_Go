package main

import "fmt"

func main() {
	type nohp string

	var NomorHP nohp = "08123456789"
	fmt.Println("Phone Number :", NomorHP)
	fmt.Println("Phone Number :", nohp("08987654321"))
}
