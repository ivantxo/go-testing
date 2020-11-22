package main

import (
	"fmt"
)

func main() {
	// boolean declaration
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	// boolean declaration 2
	var h bool
	fmt.Printf("%v, %T\n", h, h)

	// logical tests
	m := 1 == 1
	o := 1 == 2
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", o, o)
}
