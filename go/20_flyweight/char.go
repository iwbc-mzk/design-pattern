package main

import (
	"fmt"
)

// Flyweightクラス
// それぞれのオブジェクトに共通なフィールド(内因的状態 intrinsic state)を保持する
// 今回はrune文字のみだが、通常はメモリ消費が大きいものがFlyweightになる
type Char struct {
	char rune
}

// 初期化時に共通なフィールド(内因的状態 intrinsic state)を渡す。
func NewChar(char rune) *Char {
	fmt.Printf("Create character instance. [char: %s]\n", string(char))
	return &Char{char: char}
}

func (c *Char) Print() {
	fmt.Print(string(c.char))
}