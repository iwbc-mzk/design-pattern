package main

import (
	. "prototype/framework"
)

type MessageBox struct {
	meessages []string
}

func NewMessageBox() *MessageBox {
	return &MessageBox{meessages: []string{}}
}

func (m *MessageBox) Use(s string) {
	m.meessages = append(m.meessages, s)
}

func (m *MessageBox) CreateClone() Product {
	return &MessageBox{meessages: m.meessages[:]}
}
