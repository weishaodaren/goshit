package main

import "fmt"

func main()  {
	fv := func(u string) { fmt.Printf("Hello World: %T", u)}
	fv("sad")
}