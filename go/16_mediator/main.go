package main

import (
	"sync"
)

func main() {
	urls := []string{
		"https://www.yahoo.co.jp/",
		"https://www.google.com/",
		"https://github.com/",
	}
	var wg sync.WaitGroup

	manager := NewManger()

	request := func(url string) {
		wg.Add(1)
		defer wg.Done()
		r := NewRequest(url, manager)
		r.Send()
	}

	for _, url := range urls {
		go request(url)
	}
	wg.Wait()
}
