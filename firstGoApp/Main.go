package main

import (
	"fmt"
)

func main() {
	// anonymous structs
	aDoctor := struct{ name string }{name: "John Pertwee"}
	fmt.Println(aDoctor)

	// unless maps and slices, structs have independent data sets
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(anotherDoctor)

	// We could have by reference data???
	andAnother := &aDoctor
	andAnother.name = "Ivan Ramirez"
	fmt.Println(andAnother)
}
