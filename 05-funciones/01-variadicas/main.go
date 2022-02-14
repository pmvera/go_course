package main

import "fmt"

// Primero van los parámetros determinados y después los indeterminados (...)
func sumar(nombre string, numeros ...int) (string, int) {
	var total int

	mensaje := fmt.Sprintf("La suma de %s es:", nombre)
	for _, num := range numeros {
		total += num
	}

	return mensaje, total
}

func main() {
	mensaje, result := sumar("Pablo", 4, 5, 6, 10)

	fmt.Println(mensaje, result)
}
