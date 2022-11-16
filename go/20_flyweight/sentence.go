package main

import (
	"fmt"
)

// Contextクラス
// 外因的状態(extrinsic state)を保持する
// オブジェクトの数だけ生成されることになるが、メモリ消費量が大きい部分は全てFlyweightクラスに分離されているため通常は問題にならない
type Sentence struct {
	charList []*Char
}

func NewSentence(runeList []rune) *Sentence {
	charList := []*Char{}
	factory := GetCharFactory()
	for _, r := range runeList {
		char := factory.GetChar(r)
		charList = append(charList, char)
	}
	return &Sentence{charList: charList}
}

func (s *Sentence) Print(decorator string) {
	fmt.Print(decorator)
	for _, c := range s.charList {
		c.Print()
	}
	fmt.Println(decorator)
}