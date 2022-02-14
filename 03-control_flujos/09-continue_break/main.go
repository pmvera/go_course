package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 5 {
			fmt.Println("Saltado")
			continue
		}
		if i == 8 {
			fmt.Println("Salida")
			break
		}
		fmt.Println(i)
	}
}
