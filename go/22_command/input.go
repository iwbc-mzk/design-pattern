package main

import (
	"fmt"
)

type InputCommand struct {
	editor *Editor
	text string
}

func NewInputCommand(text string, editor *Editor) *InputCommand {
	return &InputCommand{
		editor: editor,
		text: text,
	}
}

func (c *InputCommand) execute() {
	fmt.Printf("<< Input command >> (%s)\n", c.text)
	c.editor.InputText(c.text)
	c.editor.Show()
}