package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	fmt.Print(hola)
	fmt.Println(mundo)

	nombre := "Pablo"
	edad := 28

	fmt.Printf("Hola, %s y %d \n", nombre, edad)
	fmt.Printf("Hola, %v y %v \n", nombre, edad) // Con %v no se sabe el tipo de datos que se va a imprimir

	mensaje := fmt.Sprintf("Hola, %s y %d", nombre, edad)

	fmt.Println(mensaje)

	fmt.Printf("nombre: %T \n", nombre) // %T para conocer el tipo de variable

	fmt.Print("Ingrese un nombre: ")
	fmt.Scanln(&nombre)
	fmt.Println("Otro nombre:", nombre)

}
