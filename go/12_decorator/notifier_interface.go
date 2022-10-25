package main

// Conponentインターフェースではデコレーターで共通して変更可能な処理を宣言
type INotifier interface {
	Send(message string)
}
