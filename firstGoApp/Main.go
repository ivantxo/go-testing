package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	// receiving go routine
	go func() {
		i := <-ch // pulling data from the channel
		fmt.Println(i)
		wg.Done()
	}()

	// sending go routine
	go func() {
		ch <- 42 // sending data to the channel
		wg.Done()
	}()
	wg.Wait()
}
