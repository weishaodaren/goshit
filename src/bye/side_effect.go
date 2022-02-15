package main

import "fmt"

func main() {
	n := 0
	multiply(10, 5, &n)
	fmt.Println("Multiply:", *&n, n)
}

func multiply(a, b int, reply *int) {
	*reply = a * b
}
