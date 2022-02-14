package main

import (
	"fmt"
)

func calcularIgv(precio float64, impuesto float64) float64 {
	return precio * (impuesto / 100)
}

func calcular_total(precio float64, igv float64) float64 {
	return precio + igv
}

func main() {
	var precio float64

	fmt.Print("Elija el precio de venta: ")
	fmt.Scanln(&precio)

	precio_impuesto := calcularIgv(precio, 18.0)

	fmt.Println("El impuesto es:", precio_impuesto,
		"€ y el precio de venta es:", calcular_total(precio, precio_impuesto), "€")
}
