package figures

import "math"

type Circle struct {
	Radious float64
}

func (c *Circle) Area() float64 {
	return math.Pi * (c.Radious * c.Radious)
}

func (c *Circle) Perimeter() float64 {
	return 2.0 * math.Pi * c.Radious
}
