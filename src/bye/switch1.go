package main

import "fmt"

func main() {
	var num1 int = 100

	switch num1 {
	case 98, 99:
		fmt.Println("It is eaual to 98")

	case 100:
		fmt.Println("It is eaual to 100")

	default:
		fmt.Println("It is nor equal to 98 or 100")
	}

}
