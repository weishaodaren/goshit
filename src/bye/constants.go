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

func main() {
	fm.Println(beef, two, c, Monday, Tuesday, Wednesday, Unknown, Female, Male)
}
