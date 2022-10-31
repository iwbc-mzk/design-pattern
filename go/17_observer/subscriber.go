package main

type Subscriber interface {
	Update(n Notification) // Publisherからの更新通知に応答する処理
}