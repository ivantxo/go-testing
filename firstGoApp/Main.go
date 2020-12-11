package main

import (
	"fmt"
)

func main() {
	greeting := "Hello"
	name := "Sonia"
	sayGreeting(&greeting, &name)
	fmt.Println(name)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ivan"
	fmt.Println(*name)
}
