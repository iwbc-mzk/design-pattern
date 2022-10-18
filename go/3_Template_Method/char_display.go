package main

import (
	"fmt"
)

type CharDisplay struct {
	ch rune
}

func NewCharDisplay(ch rune) *CharDisplay {
	c := CharDisplay{ch: ch}
	return &c
}

func (c *CharDisplay) Open() {
	fmt.Print("<<")
}

func (c *CharDisplay) Print() {
	fmt.Print(string(c.ch))
}

func (c *CharDisplay) Close() {
	fmt.Println(">>")
}