package main

import "fmt"

func main() {
	// Adding tags is a good way to break an outer loop
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 2 {
				break Loop
			}
		}
	}
}
