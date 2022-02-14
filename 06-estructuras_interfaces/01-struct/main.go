package main

import "fmt"

// Struct persona
type Persona struct {
	nombre string
	edad   int
}

// Métodos persona
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)
}

func (p *Persona) editarNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

// Herencia
type Empleado struct {
	Persona
	sueldo float64
}

func main() {
	p1 := Persona{"Pablo", 28}

	fmt.Println(p1)
	p1.nombre = "Moreno"
	fmt.Println(p1)

	p2 := Persona{
		edad:   24,
		nombre: "Pepe",
	}
	fmt.Println(p2)

	// Métodos
	p1.imprimir()
	p2.imprimir()
	p1.editarNombre("Moreno")
	p1.imprimir()

	// Herencia
	emp1 := Empleado{
		sueldo: 1550.0,
	}
	emp1.nombre = "Pedro"
	emp1.edad = 30
	emp1.imprimir()
	fmt.Println(emp1)
}
