package main

import "fmt"

type TZ int
type Rope string

func main() {
	var a, b TZ = 3, 4
	c := a + b
	fmt.Printf("c has the value %d", c)

	var d Rope = "weishaodren"
	print(d)
}
