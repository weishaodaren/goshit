package main

import "fmt"

func main() {
	p2 := Add2()
	fmt.Printf("Call Add2 for gives: %v\n", p2(3)) // 5
	TwoAdder := Adder(2)                           // func(b int) int
	fmt.Printf("The result is: %v\n", TwoAdder(3)) // 5
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
