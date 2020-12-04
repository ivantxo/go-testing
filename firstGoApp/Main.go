package main

import "fmt"

func main() {
	guess := -5
	// short circuiting: in this case guess is less than 1, then it doesn't evaluate the
	// other conditions, it just assumes is true.
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("The guest must be between 1 and 100")
	}
}

func returnTrue() bool {
	fmt.Println("Returning true")
	return true
}
