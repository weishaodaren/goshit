package main

import "fmt"

func main() {
	a := 1
	b := 2

	defer fmt.Println(a + b)

	a = 2

	fmt.Println("main")
}
