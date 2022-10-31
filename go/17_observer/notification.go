package main

type Notification struct {
	Title string
	Message string
}

func NewNotification(title, message string) *Notification {
	return &Notification{
		Title: title,
		Message: message,
	}
}