package main

import (
	"fmt"
	"time"
)

func main() {
	var msg = "Hello"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye"
	// explicitly using sleep to wait for the go routine
	time.Sleep(100 * time.Millisecond)
}
