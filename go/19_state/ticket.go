package main

import (
	"fmt"
)

// Contextオブジェクト
// 現在の状態を表すStateオブジェクトの参照を保持する。
type Ticket struct {
	title string
	state State
}

func NewTicket(title string) *Ticket {
	return &Ticket{
		title: title,
		state: GetOpneInstance(),
	}
}

func (t *Ticket) ChangeState(state State) {
	fmt.Printf("change state: '%s' -> '%s'\n", t.state.GetState(), state.GetState())
	t.state = state
}

// 状態変更リクエストの処理をStateオブジェクトに委任する。
// -> 現在の状態に適切な処理をContextで判断する必要がない。
func (t *Ticket) Approve() {
	t.state.Approve(t)
}

func (t *Ticket) Reject() {
	t.state.Reject(t)
}

func (t *Ticket) Close() {
	t.state.Close(t)
}

func (t *Ticket) ShowState() {
	fmt.Printf("<< Current State: %s >>\n", t.state.GetState())
}