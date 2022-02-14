package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

func sumar(a int, b int) int {
	return a + b
}

func main() {
	saludar("Pablo")
	sum := sumar(5, 3)
	fmt.Println("La suma es:", sum)
}
