package main

import "fmt"

func main() {
	// loop slices or arrays s := [3]string{"a", "b", "c"}
	s := []string{"a", "b", "c"}
	for k, v := range s {
		fmt.Println(k, v)
	}
	fmt.Println()

	// loop over maps
	statePopulations := map[string]int{
		"NSW": 9494,
		"QLD": 123123,
		"NT":  73694,
		"WA":  8382,
		"SA":  3737,
		"VIC": 93832,
		"TAS": 88430,
		"ACT": 83729,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}
	fmt.Println()

	// loop through a string
	f := "Hello Go!"
	for k, v := range f {
		fmt.Println(k, v) // unicode representation of each character
	}
	fmt.Println()
	for k, v := range f {
		fmt.Println(k, string(v)) // string representation of the unicode of each character
	}
}
