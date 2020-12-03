package main

import (
	"fmt"
)

func main() {
	// maps

	// built-in function
	// statePopulations := make(map[string]int)

	// literal sintax
	statePopulations := map[string]int{
		"QLD": 123,
		"NSW": 456,
		"SA":  789,
		"TAS": 012,
		"VIC": 345,
		"WA":  678,
	}
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["QLD"])
	fmt.Println(statePopulations)

	// Add new elements to map
	statePopulations["Papua"] = 455
	fmt.Println(statePopulations["Papua"])
	// The order of elements in a map is not guaranted to be the same
	fmt.Println(statePopulations)

	// delete elements from a map
	delete(statePopulations, "Papua")
	fmt.Println(statePopulations)

	// get an element with the comma, ok syntax
	pop, ok := statePopulations["Papua"]
	fmt.Println(pop, ok)

	// size of a map
	fmt.Println(len(statePopulations))

	// data is passed by reference
	sp := statePopulations
	delete(sp, "QLD")
	fmt.Println(sp)
	fmt.Println(statePopulations)
}
