package main

import (
	"fmt"
)

type MobileNotifier struct {
	name string
}

func NewMobileNotifier() *MobileNotifier {
	return &MobileNotifier{name: "Mobile"}
}

func (m *MobileNotifier) Update(n Notification) {
	fmt.Printf("[%s] %s (to %s)\n", n.Title, n.Message, m.name)
}