package main

import (
	"fmt"
)

type JpegExporter struct {}

func NewJpegExporter() *JpegExporter {
	return &JpegExporter{}
}

func (p *JpegExporter) VisitForSquare(s *Square) {
	fmt.Println("export square fugure to jpeg")
}

func (p *JpegExporter) VisitForCircle(c *Circle) {
	fmt.Println("export circle fugure to jpeg")
}

func (p *JpegExporter) VisitForRectangle(r *Rectangle) {
	fmt.Println("export rectangle fugure to jpeg")
}