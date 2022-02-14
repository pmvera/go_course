package main

import "fmt"

func main() {
	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println(numeros)
	fmt.Println(numeros[1])

	nombres := [3]string{"Pablo", "Moreno", "Vera"}
	fmt.Println(nombres)

	// Array con longitud definida por los valores
	colores := [...]string{
		"Rojo",
		"Verde",
		"Negro",
		"Azul",
	}
	fmt.Println(colores, len(colores))

	monedas := [...]string{
		0: "Dolar",
		2: "Soles",
		3: "Pesos",
		5: "Euros",
	}
	fmt.Println(monedas, len(monedas))

	submonedas := monedas[:4]
	fmt.Println(submonedas)
}
