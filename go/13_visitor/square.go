package main

type Square struct {
	side float64
}

func NewSquare(side float64) *Square {
	return &Square{side: side}
}

// acceptメソッド内でvisitorの各要素ごとに実装されたメソッドを呼び出す。
// メソッドのオーバーーロードが許される言語の場合はvisit(self)のようにできる。
func (s *Square) Accept(v Visitor) {
	// Goではメソッドのオーバーロードができないため、ここで要素ごとのメソッドを呼び出す。
	v.VisitForSquare(s)
}

func (s *Square) GetSide() float64 {
	return s.side
}
