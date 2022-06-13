package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

func main() {
	r1 := Rectangle{12.0, 6.0}
	r2 := Rectangle{8.0, 6.0}

	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", area(r2))
}
