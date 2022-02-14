package main

import (
	"fmt"
)

func suma(a int, b int) int {
	return a + b
}

func main() {
	var num1, num2 int

	fmt.Print("Elija un número: ")
	fmt.Scanln(&num1)
	fmt.Print("Elija otro número: ")
	fmt.Scanln(&num2)

	fmt.Println("La suma es:", suma(num1, num2))
}
