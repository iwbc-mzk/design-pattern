package main

import (
	"fmt"
)

var close *Close

type Close struct {}

func init() {
	close = &Close{}
}

func GetCloseInstance() *Close {
	return close
}

func (c *Close) Approve(ticket *Ticket) {
	fmt.Println("this ticket is already closed.")
}

func (c *Close) Reject(ticket *Ticket) {
	// Close -> In Progress
	ticket.ChangeState(GetInProgressInstance())
}

func (c *Close) Close(ticket *Ticket) {
	fmt.Println("this ticket is already closed.")
}

func (c *Close) GetState() string {
	return "Close"
}