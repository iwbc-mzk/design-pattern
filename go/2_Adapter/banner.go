package main

import (
	"fmt"
)
// Adaptee
type Banner struct {
	str string
}

func NewBanner(str string) *Banner {
	b := Banner{str: str}
	return &b
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("(%s)\n", b.str)
}

func (b *Banner) ShowWithAster() {
	fmt.Printf("*%s*\n", b.str)
}
