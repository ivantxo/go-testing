package main

import (
	"fmt"
	"reflect"
)

// Animal struct
type Animal struct {
	Name   string `required max: "100"`
	Origin string
}

func main() {
	// tags
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
