package main

import "fmt"

func main() {

	if nombre, edad := "Pablo", 28; nombre == "Pablo" {
		fmt.Println("Hola,", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d\n", nombre, edad)
	}

	users := make(map[string]string)

	users["Pablo"] = "pablo@gmail.com"
	users["Moreno"] = "moreno@gmail.com"

	if correo, ok := users["Pablo"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor.")
	}
}
