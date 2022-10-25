package main

import (
	"fmt"
)

// デコレーターは関数を引数として受け取る。
// 受け取った関数の実行前後に追加処理を行う無名関数を戻り値として返す。
func decorator1(f func(str string)) func(str string) {
	return func(s string) {
		fmt.Println("<< Started >>")
		f(s)
		fmt.Println("<< Done! >>")
	}
}

func decorator2(f func(str string)) func(str string) {
	return func(s string) {
		fmt.Println("== Waiting ==")
		f(s)
		fmt.Println("== Processing ==")
	}
}

func test(str string) {
	fmt.Println(str)
}

// クライアント側はComponentインターフェースを介してオブジェクトと通信する。
// そのためComponentインターフェースを実装してさえいれば、クライアント側の変更なしに
// 新しいデコレーターや具象コンポーネントを追加可能。
func main() {
	// fmt.Println("====== Send by only mail ======")
	// notifier := NewNotifier()
	// notifier.Send("不正な利用を検知しました。")

	// fmt.Println("====== Send by mail and slack ======")
	// notifierWithSlack := NewSlackDecorator(notifier)
	// notifierWithSlack.Send("本日は10時より会議の予定です。")

	// fmt.Println("====== Send by mail, slack and twitter ======")
	// notifierWithSlackAndTwitter := NewTwitterDecorator(notifierWithSlack)
	// notifierWithSlackAndTwitter.Send("今日の天気は晴れです。")

	decorator1(decorator2(test))("Hello World")
}
