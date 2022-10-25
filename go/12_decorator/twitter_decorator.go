package main

import (
	"fmt"
)

type TwitterDecorator struct {
	nd NotifierDecorator
}

func NewTwitterDecorator(n INotifier) *TwitterDecorator {
	return &TwitterDecorator{
		nd: *NewNotifierDecorator(n),
	}
}

func (d *TwitterDecorator) Send(message string) {
	fmt.Printf("[Twitter] Send message: \"%s\"\n", message)
	d.nd.Send(message)
}
