package main

import (
	"fmt"
	"paquetes/models"
)

func main() {
	p1 := models.Persona{}
	p1.Constructor("Pablo", 28)

	fmt.Println(p1)
	fmt.Println(p1.GetNombre(), p1.GetEdad())
	p1.SetNombre("Moreno")
	p1.SetEdad(30)
	fmt.Println(p1.GetNombre(), p1.GetEdad())
}
