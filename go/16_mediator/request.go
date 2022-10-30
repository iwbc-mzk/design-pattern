package main

import (
	"fmt"
	"time"
)

type Request struct {
	url      string
	mediator Mediator  // Mediatorオブジェクトを保持する。
}

func NewRequest(url string, mediator Mediator) *Request {
	return &Request{
		url:      url,
		mediator: mediator,
	}
}

func (r *Request) Send() {
	if !r.mediator.canRequest(r) {
		fmt.Printf("[url: %s] Another request is sending.\n", r.url)
		return
	}
	fmt.Printf("[url: %s] Sending request.\n", r.url)
	time.Sleep(time.Second * 1)
	fmt.Printf("[url: %s] Sending request is finished.\n", r.url)

	// 処理の完了をMediatorに通知する。
	r.mediator.notifyRequestFinished()
}

func (r *Request) permitSending() {
	fmt.Printf("[url: %s] Request sending is permitted.\n", r.url)
	r.Send()
}