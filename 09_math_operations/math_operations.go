package main

import "fmt"

func main() {
	// Addition
	var add = 10 + 10
	fmt.Println("10 + 10 = ", add)

	// Subtraction
	var a = 50
	var b = 10
	var c = a - b
	fmt.Println(a, "-", b, "=", c)

	// Multiplication
	var e = 6
	var f = 2
	var g = e * f
	fmt.Println("6 * 2 =", g)

	// Division
	var div = 96 / 6
	fmt.Println("96 / 6 =", div)

	// Modulo
	var mod = 10 % 3
	fmt.Println("10 mod 3 =", mod)

	// Augmented assignment
	var i = 5
	i += 10 // i = i + 10
	fmt.Println(i)

	// Unary Operator
	var z = 1
	z++ // z = z + 1
	fmt.Println(z)

	// Positive & Negative
	var positive = +10 // = 10
	var negative = -10

	fmt.Println(positive)
	fmt.Println(negative)
}
