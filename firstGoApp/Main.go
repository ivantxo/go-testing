package main

import (
	"fmt"
)

func main() {
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)

	var ru rune = 'a'
	fmt.Printf("%v, %T\n", ru, ru)
}
