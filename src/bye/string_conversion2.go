package main

import (
	"fmt"
	"strconv"
	"os"
)

func main(){
	var orig string = "ABC" 
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	// an, err := strconv.Atoi(orig)
	an, _ := strconv.Atoi(orig)
	// if err != nil {
	if 	_, err := strconv.Atoi(orig); err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error \n", orig)
		os.Exit(1)
		fmt.Printf("OOOOorig %s is not an integer - exiting with error \n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
