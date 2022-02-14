package main

import "fmt"

func main() {
	a := 2

	// Operadores en asignacion
	a = a + 2
	fmt.Println(a)

	// Suma en asignación
	a += 2
	fmt.Println(a)

	// Resta en asignación
	a -= 2
	fmt.Println(a)

	// Multiplicacion en asignación
	a *= 4
	fmt.Println(a)

	// Division en asignación
	a /= 2
	fmt.Println(a)

	// Modulo en asignación
	a %= 2
	fmt.Println(a)
}
