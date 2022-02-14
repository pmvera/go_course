package main

import (
	"fmt"
)

func cociente(a int, b int) int {
	return a / b
}

func resto(a int, b int) int {
	return a % b
}

func main() {
	var num1, num2 int

	fmt.Print("Elija un número: ")
	fmt.Scanln(&num1)
	fmt.Print("Elija otro número: ")
	fmt.Scanln(&num2)

	fmt.Println("El cociente es:", cociente(num1, num2),
		"y el resto:", resto(num1, num2))
}
