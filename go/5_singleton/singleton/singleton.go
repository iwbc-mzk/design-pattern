package singleton

import (
	"time"
	"fmt"
	// "sync"
	// "math/rand"
)

var instance *singleton
// var once sync.Once
// var mu sync.Mutex

type singleton struct {
	id int
}

func init() {
	fmt.Println("start initialize")
	instance = &singleton{id: 1}
	fmt.Printf("init: %p\n", instance)
}

func GetSingleton() *singleton {
	time.Sleep(1 * time.Second)
	return instance
}

// 下記ではスレッドセーフにならないためNG
// func GetSingleton() *singleton {
//     if instance == nil {
// 		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
//         instance = &singleton{id: 1}
// 		fmt.Printf("init: %p\n", instance)
//     }
//     return instance
// }

// func GetSingleton() *singleton {
// 	once.Do(func() {
// 		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
//         instance = &singleton{id: 1}
// 		fmt.Printf("init: %p\n", instance)
// 	})
//     return instance
// }

// func GetSingleton() *singleton {
// 	mu.Lock()
// 	if instance == nil {
// 		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
// 		instance = &singleton{id: 1}
// 		fmt.Printf("init: %p\n", instance)
// 	}
// 	mu.Unlock()
//     return instance
// }