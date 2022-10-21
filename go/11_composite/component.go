package main

// Leaf, Compositeに共通するインターフェースを宣言
type Component interface {
	show(indent int)
  // 場合によってはここで要素追加用メソッドを宣言する場合もある。
  // その場合、クライアント側で全ての要素を同等に扱える利点がある。
  // add(c Component)
}