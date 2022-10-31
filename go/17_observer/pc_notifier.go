package main

import (
	"fmt"
)

type PCNotifier struct {
	name string
}

func NewPCNotifier() *PCNotifier {
	return &PCNotifier{name: "PC"}
}

func (p PCNotifier) Update(n Notification) {
	fmt.Printf("[%s] %s (to %s)\n", n.Title, n.Message, p.name)
}