package main

import "fmt"

func main() {
	var first int = 10
	var cond int

	if first <= 0 {
		fmt.Printf("First is less than or equal to 0\n")
	} else if first > 0 && first < 5 {
		fmt.Printf("First is between 0 and 5\n")
	} else {
		fmt.Printf("First is 5 or greater\n")
	}

	if cond = 5; cond > 7 {
		fmt.Printf("Cond is greater than 10 \n")
	} else {
		fmt.Printf("Cond is not greater than 10\n")
	}
}
