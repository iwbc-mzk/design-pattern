package main

import (
	"fmt"
)

func main() {
	shapes := []Shape{
		NewSquare(5.6),
		NewCircle(1),
		NewRectangle(3, 2.5),
	}

	// クライアント側は具体的な要素を把握することなくVisitorの操作を実行可能。
	for _, shape := range shapes {
		areaCalculator := NewAreaCalculator()
		perimeterCalculator := NewPerimeterCalculator()
		exporter := NewJpegExporter()
		shape.Accept(areaCalculator)
		shape.Accept(perimeterCalculator)
		shape.Accept(exporter)
		fmt.Println()
	}

	// Goで直接Visitorのメソッドを呼び出す場合は、クライアント側でオブジェクトが何の構造体かを知る必要がある。
	// メソッドのオーバーロードが許されている場合でも、親要素と子要素の区別ができないため2重ディスパッチが必要になる。
	// areaCalculator.VisitForSquare(square)
}