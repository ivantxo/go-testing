package main

import (
	"fmt"
	"sync"
)

// asynchronous routine using Wait Group
var wg = sync.WaitGroup{}

func main() {
	var msg = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Goodbye"
	wg.Wait()
}
