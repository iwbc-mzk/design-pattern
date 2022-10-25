package main

import (
	"fmt"
)

// 具象コンポーネントはデフォルトの処理を実装する。
type Notifier struct{}

func NewNotifier() *Notifier {
	return &Notifier{}
}

// デフォルト処理
func (n *Notifier) Send(message string) {
	fmt.Printf("[Mail] Send message: \"%s\"\n", message)
}
