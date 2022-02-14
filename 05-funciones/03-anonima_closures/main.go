package main

import (
	"fmt"
	"strings"
)

// Closures
func repeat(n int) func(cadena string) string {

	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {

	// Funcion anonima
	func() {
		fmt.Println("Hola desde la funcion anonima")
	}()

	my_func := func(nombre string) string {
		return fmt.Sprintf("Hola %s, desde la funcion anonima asignada", nombre)
	}
	fmt.Println(my_func("Pablo"))

	repeat3 := repeat(3)
	fmt.Println(repeat3("Hola"))
	fmt.Println(repeat3("Mundo"))

	repeat5 := repeat(5)
	fmt.Println(repeat5("Pablo"))
	fmt.Println(repeat5("Moreno"))

}
