package main

import (
	"paquetes/figures"
)

func main() {
	s := figures.Square{Side: 8.0}
	c := figures.Circle{Radious: 3.0}
	figures.Measures(&s)
	figures.Measures(&c)

}
