package main

import "fmt"

func main() {
	i := 10
	switch { // <== tagless switch statement
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // it will fall through the next case

	case i <= 20:
		fmt.Println("less than or equal to twenty")

	default:
		fmt.Println("another number")
	}
}
