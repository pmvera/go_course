package main

import "fmt"

func main() {
	nombres := [...]string{"Pablo", "Moreno", "Pepe"}

	// for i := 0; i < len(nombres); i++ {
	// 	fmt.Println(nombres[i])
	// }

	for indice, elemento := range nombres {
		fmt.Println(indice, elemento)
	}
}
