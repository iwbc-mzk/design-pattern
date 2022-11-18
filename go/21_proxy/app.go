package main

import "fmt"

type App struct {}

func NewApp() *App {
	return &App{}
}

func (a *App) HandleRequest(url, method string, body map[string]string) (int, string) {
	fmt.Println("App received a request.")

	if url == "/users" && method == "GET" {
		return 200, "['tanaka', 'kato', 'yamamoto']"
	}

	return 404, "Not Found"
}