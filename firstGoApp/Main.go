package main

import (
	"fmt"
	"sync"
)

// asynchronous routine using Wait Group
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	// This will print Hello <counter>, but the counter will be inconsistent 1, 2, 1, 1, 10, 9, 9
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

// func sayHello() {
// 	fmt.Println("Hello #", counter)
// 	wg.Done()
// }

func sayHello() {
	fmt.Println("Hello #", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
