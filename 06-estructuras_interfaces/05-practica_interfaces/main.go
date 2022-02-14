package main

import (
	"fmt"
	"math"
)

type Figure interface {
	area() float64
	perimeter() float64
}

type Square struct {
	side float64
}

type Circle struct {
	radious float64
}

func (s *Square) area() float64 {
	return s.side * s.side
}

func (s *Square) perimeter() float64 {
	return 4.0 * s.side
}

func (c *Circle) area() float64 {
	return math.Pi * (c.radious * c.radious)
}

func (c *Circle) perimeter() float64 {
	return 2.0 * math.Pi * c.radious
}

func area(figure Figure) float64 {
	return figure.area()
}

func perimeter(figure Figure) float64 {
	return figure.perimeter()
}

func main() {
	s := Square{
		side: 2.0,
	}
	c := Circle{
		radious: 4.0,
	}

	fmt.Println("Area cuadrado:", area(&s))
	fmt.Println("Area círculo:", area(&c))
	fmt.Println("Perímetro cuadrado:", perimeter(&s))
	fmt.Println("Perímetro círculo:", perimeter(&c))
}
