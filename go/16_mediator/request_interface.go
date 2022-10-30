package main

type IRequest interface {
	Send()
	permitSending()  // Mediatorからリクエスト許可をもらうためのメソッド
}