package main

import (
	"fmt"
)

func main() {
	n := 42
	fmt.Printf("%v, %T\n", n, n)

	// Integer types
	// int8, int16, int32, int64
	// uint8, uint16, uint32

	var a int = 3
	var b int8 = 10
	// fmt.Println(a + b) // Type mismatch
	fmt.Println(a + int(b))
}
