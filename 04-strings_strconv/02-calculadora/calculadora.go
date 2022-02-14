package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {
	var result int

	valores := strings.Split(expresion, "+")
	fmt.Printf("%T\n", valores)

	for _, valor := range valores {
		num, err := strconv.Atoi(valor)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Error: expresión inválida")
		} else {
			result += num
		}
	}

	return result
}

func main() {
	var expresion string
	var result int

	fmt.Print("=>")
	fmt.Scanln(&expresion)

	result = sumar(expresion)

	fmt.Printf("=> %d\n", result)
}
