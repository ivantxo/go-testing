package main

import (
	"fmt"
)

func main() {
	g := greeter{
		greeting: "Hola",
		name:     "Ivan",
	}
	g.greet()
	fmt.Println("The new name is:", g.name)
}

type greeter struct {
	greeting string
	name     string
}

// If using * the values of the original struct can be manipulated
// func (g *greeter) greet() {
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}
