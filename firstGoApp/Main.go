package main

import (
	"fmt"
)

func main() {
	// this will print start, this was deferred, panic: something bad happened
	// the order of execution is panic after defer
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end")
}
