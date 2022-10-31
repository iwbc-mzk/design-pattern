package main

import (
	"math/rand"
	"time"
	"fmt"
)

type EarthquakeSensor struct {
	subscribers []Subscriber
}

func NewEarthquakeSensor() *EarthquakeSensor {
	return &EarthquakeSensor{}
}

// Subscriberを引数として受け取り、購読登録処理を行う。
// 関数を引数として渡せる場合は、構造体ではなく関数を引数として実装することも可能。
func (s *EarthquakeSensor) Subscribe(sub Subscriber) {
	s.subscribers = append(s.subscribers, sub)
}

func (s *EarthquakeSensor) Unsubscribe(sub Subscriber) {
	for i, so := range s.subscribers {
		if so == sub {
			s.subscribers = append(s.subscribers[:i], s.subscribers[i+1:]...)
		}
	}
}

func (s *EarthquakeSensor) NotifiyAll() {
	rand.Seed(time.Now().UnixNano())
	t := time.Date(2022, time.Month(rand.Intn(12)), rand.Intn(29), rand.Intn(24), rand.Intn(60), rand.Intn(60), 0, time.UTC)
	notification := NewNotification(
		"Earthquake Alert",
		fmt.Sprintf("The earthquake occurred at %s", t.Format("2006-01-02 15:4:5")),
	)
	for _, o := range s.subscribers {
		o.Update(*notification)
	}
}