package main

import "fmt"

func main() {
	var a = 1
	var b = 2

	defer func(a, b int) {
		fmt.Println(a + b)
	}(a, b)

	a = 2

	fmt.Println("main")
}
