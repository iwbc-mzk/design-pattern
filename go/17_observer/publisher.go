package main

type Publisher interface {
	Subscribe(s Subscriber)
	Unsubscribe(s Subscriber)
	NotifiyAll()
}