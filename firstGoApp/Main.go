package main

import (
	"fmt"
)

func main() {
	// Removing elements from slices
	a := []int{1, 2, 3, 4, 5}

	b := a[1:] // removes the first element
	fmt.Println(b)

	c := a[:len(a)-1] // removes an element at the end of slice
	fmt.Println(c)

	d := append(a[:2], a[3:]...) // removes en element from the middle
	fmt.Println(d)

	fmt.Println(a) // becareful when manipulating arrays, the original array might have changed
}
