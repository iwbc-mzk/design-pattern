package main

type Circle struct {
	radius int
}

func NewCircle(radius int) *Circle {
	return &Circle{radius: radius}
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}

func (c *Circle) GetRadius() int {
	return c.radius
}
