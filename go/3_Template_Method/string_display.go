package main

import (
	"fmt"
)

type StringDisplay struct {
	str string
}

func NewStringDisplay(str string) *StringDisplay {
	s := StringDisplay{str: str}
	return &s
}

func (s *StringDisplay) Open() {
	fmt.Print("=")
}

func (s *StringDisplay) Print() {
	fmt.Print(s.str)
}

func (s *StringDisplay) Close() {
	fmt.Println("=")
}