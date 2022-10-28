package main

import "time"

type InputData struct {
	Id int
	RegisteredAt time.Time
	Password string
}

func NewInputData(id int, registeredAt time.Time, password string) *InputData {
	return &InputData{
		Id: id,
		RegisteredAt: registeredAt,
		Password: password,
	}
}