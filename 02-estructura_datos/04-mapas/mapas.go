package main

import "fmt"

func main() {
	dias := make(map[int]string)
	fmt.Println(dias, len(dias))

	// Agregar datos
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"
	fmt.Println(dias, len(dias))

	dias[7] = "SABADO"
	fmt.Println(dias, len(dias))

	delete(dias, 1)
	fmt.Println(dias, len(dias))

	estudiantes := make(map[string][]int)

	estudiantes["Pablo"] = []int{13, 15, 16}
	estudiantes["Vida"] = []int{14, 13, 17}
	fmt.Println(estudiantes, len(estudiantes))

	fmt.Println(estudiantes["Vida"][1:])

}
