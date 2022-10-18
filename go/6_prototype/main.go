package main

import (
	"fmt"
	. "prototype/framework"
)

func main() {
	var pr Registry = *NewRegistry()
	var mb MessageBox = *NewMessageBox()

	mb.Use("message 1")
	mb.Use("message 2")
	mb.Use("message 3")

	const name string = "messageBox"
	pr.Register(name, &mb)

	copy, _ := pr.Create(name)
	c := *(*copy).(*MessageBox)
	fmt.Printf("Original address: %p value: %v\n", &mb, mb)
	fmt.Printf("Copy address: %p value: %v\n", &c, c)
}