package main

// ElementインターフェースはVisitorインターフェースを引数としたacceptメソッドを宣言する。
type Shape interface {
	Accept(v Visitor)
}