package main

type Handler interface {
	Validate(*InputData) bool
	SetNext(Handler) Handler
}