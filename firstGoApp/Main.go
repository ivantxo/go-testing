package main

import (
	"fmt"
)

const a int16 = 27

func main() {
	const myConst int = 42
	// myConst = 27 // Error, cannot assign to const
	fmt.Printf("%v, %T\n", myConst, myConst)

	// shadowing
	const a int = 14
	fmt.Printf("%v, %T\n", a, a)
}
