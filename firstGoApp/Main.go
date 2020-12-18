package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Threads\n", runtime.GOMAXPROCS(-1))
}
