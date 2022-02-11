package main

import (
	"fmt"
	"strings"
)

func main(){
	var str string = "This is an exmaple of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}
