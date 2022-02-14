package main

import (
	"figures"

	"github.com/donvito/hellomod"
)

// Para poder usar este paquete hay que ejecutar 'go get github.com/donvito/hellomod'

func main() {
	hellomod.SayHello()

	s := figures.Square{Side: 10.0}
	figures.Measures(&s)
}
