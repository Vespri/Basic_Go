package main

import "fmt"

func main() {
	var nilaiUjian = 90
	var nilaiTugas = 75

	var lulusUjian bool = nilaiUjian > 75
	var lulusTugas bool = nilaiTugas > 75

	// checking
	fmt.Println("Lulus nilai Ujian =", lulusTugas)
	fmt.Println("Lulus nilai Tugas =", lulusUjian)

	// && (and)
	fmt.Println("And (", lulusTugas, "&&", lulusUjian, ") =", lulusTugas && lulusUjian)
	// || (or)
	fmt.Println("Or (", lulusTugas, "||", lulusUjian, ") =", lulusTugas || lulusUjian)
}
