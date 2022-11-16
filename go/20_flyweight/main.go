package main

import "fmt"

func main() {
	s1 := NewSentence([]rune{'し', 'ん', 'ぶ', 'ん', 'し'})
	s1.Print("==")

	s2 := NewSentence([]rune{'し', 'し', 'お', 'ど', 'し'})
	s2.Print("~~")

	fmt.Printf("pool size = %d", len(GetCharFactory().pool))
}