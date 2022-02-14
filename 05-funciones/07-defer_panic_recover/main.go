package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Ups! El programa no finalizó de manera correcta.")
		}
	}()

	if file, err := os.Open("holas.txt"); err != nil {
		panic("Error al abrir el fichero")
	} else {
		defer func() {
			fmt.Println("Cerrando fichero")
			file.Close()
		}()

		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)

		contenido_fichero := string(contenido[:long])
		fmt.Println(contenido_fichero)
	}
}
