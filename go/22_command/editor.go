package main

import (
	"fmt"
)

// Receiver
type Editor struct {
	text string
}

func NewEditor() *Editor {
	return &Editor{}
}

func (e *Editor) InputText(text string) {
	e.text = text
}

func (e *Editor) Show() {
	fmt.Printf("Editor's text: %s\n", e.text)
}