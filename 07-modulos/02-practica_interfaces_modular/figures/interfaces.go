package figures

import "fmt"

type figure interface {
	Area() float64
	Perimeter() float64
}

func Area(figure figure) float64 {
	return figure.Area()
}

func Perimeter(figure figure) float64 {
	return figure.Perimeter()
}

func Measures(f figure) {
	fmt.Println("Area:", f.Area())
	fmt.Println("Perimeter:", f.Perimeter())
}
