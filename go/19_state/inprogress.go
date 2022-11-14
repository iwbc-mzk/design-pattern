package main

var inProgress *InProgress

type InProgress struct {}

func init() {
	inProgress = &InProgress{}
}

func GetInProgressInstance() *InProgress {
	return inProgress
}

func (i *InProgress) Approve(ticket *Ticket) {
	// In Progress -> Close
	ticket.ChangeState(GetCloseInstance())
}

func (i *InProgress) Reject(ticket *Ticket) {
	// In Progress -> Open
	ticket.ChangeState(GetOpneInstance())
}

func (i *InProgress) Close(ticket *Ticket) {
	// In Progress -> Close
	ticket.ChangeState(GetCloseInstance())
}

func (i *InProgress) GetState() string {
	return "In Progress"
}