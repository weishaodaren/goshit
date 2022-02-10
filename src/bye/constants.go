package main

import fm "fmt"

const beef, two, c = "eat", 2, "veg"
const (
	Monday, Tuesday, Wednesday = 1, 2, 3
)
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const (
	a = iota
	b = iota
	d = iota
)

func main() {
	fm.Println(beef, two, c, Monday, Tuesday, Wednesday, Unknown, Female, Male)
	fm.Println(a, b, d)
}
