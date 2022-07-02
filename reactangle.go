package golang_united_school_homework

type Rectangle struct {
	Height, Weight float64
}

func (rectangle Rectangle) CalcPerimeter() float64 {
	return (2 * rectangle.Height) + (2 * rectangle.Weight)
}

func (rectangle Rectangle) CalcArea() float64 {
	return rectangle.Height * rectangle.Weight
}
