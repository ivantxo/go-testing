package main

import (
	"fmt"
)

func main() {
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2])
	// s[2] = "u" // this is wrong

	// string concatenation
	s2 := "this is also a string"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	// collection of bytes
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}
