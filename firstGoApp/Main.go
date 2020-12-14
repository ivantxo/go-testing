package main

import (
	"fmt"
)

func main() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
	fmt.Println("")
}

// Incrementer interface
type Incrementer interface {
	Increment() int
}

// IntCounter int
type IntCounter int

// Increment implementation
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
