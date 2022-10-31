package main

import (
	"math/rand"
	"time"
	"fmt"
)

type EarthquakeSensorForFunction struct {
	functions []func(n Notification)
}

func NewEarthquakeSensorForFunction() *EarthquakeSensorForFunction {
	return &EarthquakeSensorForFunction{}
}

func (s *EarthquakeSensorForFunction) Subscribe(f func(n Notification)) (id int) {
	id = len(s.functions) 
	s.functions = append(s.functions, f)
	return id
}

func (s *EarthquakeSensorForFunction) Unsubscribe(id int) {
	if id < len(s.functions) {
		s.functions = append(s.functions[:id], s.functions[id+1:]...)
	}
}

func (s *EarthquakeSensorForFunction) NotifiyAll() {
	rand.Seed(time.Now().UnixNano())
	t := time.Date(2022, time.Month(rand.Intn(12)), rand.Intn(29), rand.Intn(24), rand.Intn(60), rand.Intn(60), 0, time.UTC)
	notification := NewNotification(
		"Earthquake Alert",
		fmt.Sprintf("The earthquake occurred at %s", t.Format("2006-01-02 15:4:5")),
	)
	for _, f := range s.functions {
		f(*notification)
	}
}