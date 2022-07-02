package golang_united_school_homework

import "errors"

type box struct {
	shapes         []Shape
	shapesCapacity int
}

func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity == len(b.shapes) {
		return errors.New("Full capacity!")
	}

	b.shapes = append(b.shapes, shape)

	return nil
}

func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("Index out of range")
	}

	for index, value := range b.shapes {
		if i == index {
			return value, nil
		}
	}

	return nil, errors.New("No shape with this index")
}

func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, err := b.GetByIndex(i)

	if err != nil {
		return nil, err
	}

	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return shape, nil
}

func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	_, err := b.GetByIndex(i)

	if err != nil {
		return nil, err
	}

	result := make([]Shape, len(b.shapes))

	copy(result, b.shapes)

	b.shapes[i] = shape

	return result[i], nil
}

func (b *box) SumPerimeter() float64 {
	var result float64

	for _, value := range b.shapes {
		result += value.CalcPerimeter()
	}

	return result
}

func (b *box) SumArea() float64 {
	var result float64

	for _, value := range b.shapes {
		result += value.CalcArea()
	}

	return result
}

func (b *box) RemoveAllCircles() error {
	var isExist bool

	for i := 0; i < len(b.shapes); i++ {
		_, ok := b.shapes[i].(*Circle)

		if ok {
			b.ExtractByIndex(i)
			isExist = true
			i--
		}
	}

	if !isExist {
		return errors.New("Found no ircles")
	}

	return nil
}
