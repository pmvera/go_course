package main

import "fmt"

func main() {
	/*
		Descuento 10%: consumo <= 100
		Descuento 20%: consumo <= 200
		Descuento 30%: consumo > 200
		IGV: 19%
	*/
	var consumo, descuento float64
	var datosDescuento string

	fmt.Print("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		if consumo <= 100 {
			// Descuento 10%
			datosDescuento = "10%"
			descuento = consumo * 0.10
		} else if consumo > 100 && consumo <= 200 {
			// Descuento 20%
			datosDescuento = "20%"
			descuento = consumo * 0.20
		} else if consumo > 200 {
			// Descuento 30%
			datosDescuento = "30%"
			descuento = consumo * 0.30
		}
	} else {
		fmt.Println("Error al ingresar el consumo.")
	}

	// Operaciones
	montoDescuento := consumo - descuento
	igv := montoDescuento * 0.19
	totalPagar := montoDescuento + igv

	fmt.Println("------- FACTURA DE CONSUMO -------")
	fmt.Println("Descuento aplicado:", datosDescuento)
	fmt.Println("Consumo:", consumo)
	fmt.Println("Descuento:", descuento)
	fmt.Println("Total con descuento:", montoDescuento)
	fmt.Println("IGV:", igv)
	fmt.Println("Total:", totalPagar)

}
