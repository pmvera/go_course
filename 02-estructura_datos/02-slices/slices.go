package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3} // Slice
	fmt.Println(numeros, len(numeros))

	// Agregar datos
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)
	fmt.Println(numeros, len(numeros))

	//Sub slice
	subSlice := numeros[0:2]
	numeros[0] = 100

	fmt.Println(subSlice)
	fmt.Println(numeros)

	// Punteros: Inicio y fin [x:x]
	// Longitud: len
	// Cpacidad
	meses := []string{"Enero", "Febrero", "Marzo"}
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)
}
