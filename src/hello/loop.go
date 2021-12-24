package main

import (
	"fmt"
)

func main() {
	person := [3]string{"Tom", "Alice", "John"}
	fmt.Printf("len=%d cap=%d array=%v\n", len(person), cap(person), person)

	for k, v := range person {
		fmt.Printf("person[%d]: %s\n", k, v)
	}

	for i := range person {
		fmt.Printf("person[%d]: %s\n", i, person)
	}

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}

	for _, name := range person {
		fmt.Println("name:", name)
	}
}
