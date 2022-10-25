package main

import (
	"math"
	"fmt"
)

// 具象Visitorは要素別のメソッドを定義する。
type AreaCalculator struct {}

func NewAreaCalculator() *AreaCalculator {
	return &AreaCalculator{}
}

func (a *AreaCalculator) VisitForSquare(s *Square) {
	sa := math.Pow(float64(s.GetSide()) , 2)
	fmt.Printf("square area: %.2f\n", sa)
}

func (a *AreaCalculator) VisitForCircle(c *Circle) {
	ca :=  math.Pow(float64(c.GetRadius()) , 2) * math.Pi
	fmt.Printf("circle area: %.2f\n", ca)
}

func (a *AreaCalculator) VisitForRectangle(r *Rectangle) {
	ra := r.GetShortSide() * r.GetLongSide()
	fmt.Printf("rectangle area: %.2f\n", ra)
}