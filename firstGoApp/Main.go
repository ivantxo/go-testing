package main

import "fmt"

func main() {
	// Control flow, defer, panic and recover
	fmt.Println("start")
	defer fmt.Println("middle") // prints at the end but before the main function returns
	fmt.Println("end")
	fmt.Println()

	// prints end, middle, start lifo order
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}
