package main

import (
	"fmt"
)

func main() {
	// slices share the same memory space
	a := []int{1, 2, 3}
	b := a
	b[1] = 7
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("")

	// more on slices
	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := c[:]   // slice of all elements
	e := c[3:]  // slice from 4th element to end
	f := c[:6]  // slice of first 6 elements
	g := c[3:6] // slice the 4th, 5th and 6th elements
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println("")

	// another way to initialise slices
	h := make([]int, 3, 100)
	fmt.Println(h)
	fmt.Printf("Length: %v\n", len(h))
	fmt.Printf("Capacity: %v\n", cap(h))
	fmt.Println("")

	// capacity doubles in arrays when appending
	i := []int{}
	fmt.Println(i)
	fmt.Printf("Length: %v\n", len(i))
	fmt.Printf("Capacity: %v\n", cap(i))
	i = append(i, 1)
	fmt.Println(i)
	fmt.Printf("Length: %v\n", len(i))
	fmt.Printf("Capacity: %v\n", cap(i))
	i = append(i, 2, 3, 4, 5)
	fmt.Println(i)
	fmt.Printf("Length: %v\n", len(i))
	fmt.Printf("Capacity: %v\n", cap(i))
}
