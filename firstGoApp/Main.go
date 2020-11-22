package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var k string
	// k = string(i) // wrong type convertion to string
	k = strconv.Itoa(i) // right way of type converting
	fmt.Printf("%v, %T\n", k, k)
}
