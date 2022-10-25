package main

// Conponentインターフェースで宣言されたAPIを実装する。
type NotifierDecorator struct {
	notifier INotifier
}

func NewNotifierDecorator(n INotifier) *NotifierDecorator {
	return &NotifierDecorator{notifier: n}
}

// 基底デコレーターはラップされたコンポーネントオブジェクトにすべての処理を委任する。
// 追加の振る舞いは具象デコレーターで定義する。
func (d *NotifierDecorator) Send(message string) {
	d.notifier.Send(message)
}
