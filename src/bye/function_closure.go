package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Println(f(1), " - ")  // 1
	fmt.Println(f(20), " - ") // 21
	fmt.Println(f(300))       // 321
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		fmt.Println("delta =>", delta)
		x += delta
		fmt.Println("Plus delta =>", delta)
		return x
	}
}
