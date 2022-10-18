package main

import (
	"fmt"
	"sync"
	. "singleton/singleton"
)

func main() {
	fmt.Println("start main")

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			obj := GetSingleton()
			fmt.Printf("get%d: %p\n", i, obj)
		}(i)
	}

	wg.Wait()
}