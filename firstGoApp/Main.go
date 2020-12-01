package main

import (
	"fmt"
)

func main() {
	grades := [3]int{95, 97, 98}    // array of size 3
	grades2 := [...]int{95, 97, 98} // array which size is just the literal size {95, 97, 98}
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades: %v\n", grades2)

	var students [3]string
	students[0] = "Sonia"
	students[1] = "Ivan"
	students[2] = "Judith"
	fmt.Printf("Grades: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[0])
	fmt.Printf("Number of Students: %v\n", len(students))

	// array of arrays
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[0] = [3]int{1, 0, 0}
	fmt.Printf("Identy Matrix: %v\n", identityMatrix)

	// arrays don't share the same memory space
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("")

	// unless using pointers
	c := &a
	c[1] = 7
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
