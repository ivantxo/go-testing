package main

import (
	"fmt"
)

func main() {
	// a gets a new value, however b doesn't change as they are different memory locations
	a := 42
	b := a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)

	fmt.Println()
	var a1 int = 42
	var b1 *int = &a1
	fmt.Println(&a1, b1) // here we get the same memory address
	fmt.Println(a1, *b1) // * by de-referencing we get the value which that memory address is pointing to
	a1 = 37
	fmt.Println(a1, *b1) // both memory spaces point to the same memory location
	*b1 = 14
	fmt.Println(a1, *b1) // both values change
	fmt.Println("")

	// This will print &{42}, it an address of an object that has a field with a value of 42
	// var ms *myStruct
	// ms = &myStruct{foo: 42}
	// fmt.Println(ms)

	var ms *myStruct
	ms = new(myStruct)
	ms.foo = 42         // (*ms).foo this will also work
	fmt.Println(ms.foo) // (*ms).foo this will also work
}

type myStruct struct {
	foo int
}
