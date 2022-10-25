package main

import (
	"fmt"
)

type SlackDecorator struct {
	nd NotifierDecorator
}

func NewSlackDecorator(n INotifier) *SlackDecorator {
	return &SlackDecorator{
		nd: *NewNotifierDecorator(n),
	}
}

// ラップされたオブジェクトの処理を呼び出す。
// 実行前後に追加の処理を実行したり、処理を実行したデータを引数としてラップされたオブジェクトの処理に渡すことができる。
func (d *SlackDecorator) Send(message string) {
	fmt.Printf("[Slack] Send message: \"%s\"\n", message)
	d.nd.Send(message)
}
