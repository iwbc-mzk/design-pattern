package main

type Visitor interface {
	// 各要素ごとの処理を宣言する。
	// メソッドのオーバーロードが許されている言語の場合は、visit(s Square), vist(c Circle)...のように宣言できる。
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForRectangle(*Rectangle)
}