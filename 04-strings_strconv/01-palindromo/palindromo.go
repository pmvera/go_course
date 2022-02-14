package main

import (
	"fmt"
	"strings"
)

func reverse(cadena string) string {

	array_cadena := strings.Split(cadena, "")
	array_invertido := make([]string, 0)

	for i := len(array_cadena) - 1; i >= 0; i-- {
		array_invertido = append(array_invertido, array_cadena[i])
	}

	return strings.Join(array_invertido, "")
}

func esPalindromo(palabra string) bool {
	palabra = strings.ToLower(palabra)
	palabra = strings.ReplaceAll(palabra, " ", "")

	palabra_invertida := reverse(palabra)

	return palabra == palabra_invertida
}

func main() {
	fmt.Println("Es palíndromo:", esPalindromo("Luz azul"))

}
