package main

import (
	"fmt"
)

type ClearCommand struct {
	editor *Editor
}

func NewClearCommand(editor *Editor) *ClearCommand {
	return &ClearCommand{editor: editor}
}

func (c *ClearCommand) execute() {
	fmt.Println("<< Clear command >>")
	c.editor.InputText("")
	c.editor.Show()
}