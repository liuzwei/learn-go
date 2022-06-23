package test1

type Triangle struct {
	A, B, C float32
}

func (t Triangle) Perimeter() float32 {

	return t.A + t.B + t.C
}

func (t Triangle) Area() float32 {

	return t.A * t.B * t.C
}

type Rectangle struct {
	A, B float32
}

func (t Rectangle) Perimeter() float32 {

	return (t.A + t.B) * 2
}

func (t Rectangle) Area() float32 {

	return t.A * t.B
}
