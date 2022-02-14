package main

import "fmt"

func modificarValor(cadena *string) {
	fmt.Printf("%p\n", cadena)
	*cadena = "Hola desde la funcion"
}

func main() {
	cadena := "Hola mundo de Go"
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de la funcion:", cadena)

	// modificarValor(cadena) // Copia
	//fmt.Println("Después de la función:", cadena)

	modificarValor(&cadena) // Referencia
	fmt.Println("Después de la funcion por referencia:", cadena)
}
