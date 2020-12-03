package main

import (
	"fmt"
)

// Doctor structure
type Doctor struct {
	number     int
	actorName  string
	episodes   []string
	companions []string
}

func main() {
	// structs

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)

	fmt.Println(aDoctor.companions)
}
