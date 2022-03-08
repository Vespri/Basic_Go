package main

import "fmt"

func main() {
	// Integer conversion
	var int32 int32 = 129
	var int64 int64 = int64(int32)
	var int8 int8 = int8(int32)

	fmt.Println(int32, "in int32 =", int32)
	fmt.Println(int32, "in int64 =", int64)
	fmt.Println(int32, "in int8 =", int8)

	// String conversion
	var name = "Kresna"
	var e = name[0]
	var eString string = string(e)

	fmt.Println(name, "in string :", name)
	fmt.Println(name, "in byte(name[0]) :", e)
	fmt.Println(e, "in string :", eString)
}
