package main

import (
	"fmt"
	"log"
)

func main() {
	// panic and recover from panic
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("something went bad")
	fmt.Println("done panicking")
}
