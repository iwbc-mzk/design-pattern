package main

type Mediator interface {
	canRequest(IRequest) bool
	notifyRequestFinished() // Component役からMediator役へ処理の完了を通知するためのメソッド。
}
