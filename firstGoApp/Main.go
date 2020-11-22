package main

import (
	"fmt"
)

var i int = 42

func main() {
	fmt.Println(i)

	// Is possible here, different scopes, shadowing
	var i int = 27
	fmt.Println(i)

	// Can't re-declare variables in the same scope
	// i := 13
}
