package main

type Server interface {
	HandleRequest(url, method string, body map[string]string) (int, string)
}