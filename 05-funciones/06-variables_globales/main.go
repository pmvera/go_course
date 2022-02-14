package main

import "fmt"

var mensaje string

func funcion1() {
	mensaje = "Hola desde la funcion1"
	fmt.Println(mensaje)
}

func funcion2() {
	mensaje = "Hola desde la funcion2"
	fmt.Println(mensaje)
}

func main() {
	mensaje = "Hola desde la funciona principal"
	fmt.Println(mensaje)

	defer funcion1()
	funcion2()

	fmt.Println(mensaje)

}
