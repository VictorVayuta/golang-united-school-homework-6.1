package golang_united_school_homework

import "math"

type Circle struct {
	Radius float64
}

func (circle Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * circle.Radius
}

func (circle Circle) CalcArea() float64 {
	return math.Pi * (circle.Radius * circle.Radius)
}
