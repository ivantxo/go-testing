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
	fmt.Println("")

	// arrays don't work by reference
	aarr := [3]int{1, 2, 3}
	barr := aarr
	fmt.Println(aarr, barr)
	aarr[1] = 42 // if I change [1] on variable aarr, it doesn't change on barr
	fmt.Println(aarr, barr)
	fmt.Println("")

	// slices work by reference
	aarr1 := []int{1, 2, 3}
	barr1 := aarr1
	fmt.Println(aarr1, barr1)
	aarr1[1] = 42 // if I change [1] on variable aarr1, it will change on barr1
	fmt.Println(aarr1, barr1)
	fmt.Println("")

	// maps also work by reference
	amap := map[string]string{"foo": "bar", "baz": "buzz"}
	bmap := amap
	fmt.Println(amap, bmap)
	amap["foo"] = "qux" // if I change index "foo" of amap, it will change index "foo" on bmap
	fmt.Println(amap, bmap)
}

type myStruct struct {
	foo int
}
