package figures

type Square struct {
	Side float64
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}

func (s *Square) Perimeter() float64 {
	return 4.0 * s.Side
}
