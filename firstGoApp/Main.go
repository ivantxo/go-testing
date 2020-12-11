package main

import (
	"fmt"
)

func main() {
	sum(1, 2, 3, 4, 5)
}

// function with variadic parameters
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is: ", result)
}
