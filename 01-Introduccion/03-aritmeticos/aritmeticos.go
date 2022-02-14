package main

import "fmt"

func main() {
	a := 20
	b := 10

	// Suma
	result := a + b
	fmt.Println("Suma:", result)

	/* Resta */
	result = a - b
	fmt.Println("Resta:", result)

	// Multiplicacion
	result = a * b
	fmt.Println("Multiplicacion:", result)

	// Division
	result = a / b
	fmt.Println("Division:", result)

	// Division float
	var div float64 = 3.0 / 2.0
	fmt.Println("Division:", div)

	// Modulo
	result = 3 % 2
	fmt.Println("Division:", result)
}
