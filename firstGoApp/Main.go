package main

import (
	"fmt"
)

func main() {
	r, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
