package main

import (
	"fmt"
)

// Animal struct
type Animal struct {
	Name   string
	Origin string
}

// Bird struct
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// composition
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
}
