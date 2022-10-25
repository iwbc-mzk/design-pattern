package main

import (
	"math"
	"fmt"
)

type PerimeterCalculator struct {}

func NewPerimeterCalculator() *PerimeterCalculator {
	return &PerimeterCalculator{}
}

func (p *PerimeterCalculator) VisitForSquare(s *Square) {
	sp := s.GetSide() * 4
	fmt.Printf("square perimeter: %.2f\n", sp)
}

func (p *PerimeterCalculator) VisitForCircle(c *Circle) {
	cv := float64(c.GetRadius()) * 2 * math.Pi
	fmt.Printf("circle perimeter: %.2f\n", cv)
}

func (p *PerimeterCalculator) VisitForRectangle(r *Rectangle) {
	rv := (r.GetShortSide() + r.GetLongSide()) * 2
	fmt.Printf("rectangle perimeter: %.2f\n", rv)
}