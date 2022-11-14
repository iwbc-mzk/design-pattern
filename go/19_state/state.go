package main

// 全てのStateオブジェクトが実装すべきメソッドを宣言する。
type State interface {
	Approve(ticket *Ticket)
	Reject(ticket *Ticket)
	Close(ticket *Ticket)
	GetState() string
}