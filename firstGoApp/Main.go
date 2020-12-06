package main

import (
	"fmt"
)

func main() {
	// it will print start, defer takes the parameters when it is called and not when it is executed
	a := "start"
	defer fmt.Println(a)
	a = "end"
}
