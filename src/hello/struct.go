package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Name = "Tom"
	p1.Age = 39
	fmt.Println("p1 = ", p1)

	var p2 = Person{Name: "Burke", Age: 31}
	fmt.Println("p2 = ", p2)

	p3 := Person{Name: "WEISHAODAREN", Age: 99}
	fmt.Println("p3 = ", p3)

	p4 := struct {
		Name string
		Age int
	}{Name: "Shit", Age: 3}
	fmt.Println("p4 = ", p4)
}
