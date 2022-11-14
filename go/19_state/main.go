package main

import "fmt"

func main() {
	ticket := NewTicket("Test Ticket")

	// Open
	ticket.ShowState()
	fmt.Println()

	// Open -> In Progress
	ticket.Approve()
	ticket.ShowState()
	fmt.Println()

	// In Progress -> Close
	ticket.Approve()
	ticket.ShowState()
	fmt.Println()

	// Close -> Open
	ticket.Reject()
	ticket.Reject()
	ticket.ShowState()
	fmt.Println()

	// Open -> Close
	ticket.Close()
	ticket.ShowState()
	fmt.Println()

	// already closed
	ticket.Approve()
	ticket.ShowState()
}