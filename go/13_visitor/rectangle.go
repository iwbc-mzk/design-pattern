package main

type Rectangle struct {
	square Square
	ratio float64
}

func NewRectangle(shortSide float64, ratio float64) *Rectangle {
	return &Rectangle{
		square: *NewSquare(shortSide),
		ratio: ratio,
	}
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(r)
}

func (r *Rectangle) GetShortSide() float64 {
	return r.square.side
}

func (r *Rectangle) GetLongSide() float64 {
	return r.square.side * r.ratio
}
