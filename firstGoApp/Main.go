package main

import "fmt"

func main() {
	// type switch
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is int")

	case float64:
		fmt.Println("i is float64")

	case string:
		fmt.Println("i is string")

	default:
		fmt.Println("i is another type")
	}
}
