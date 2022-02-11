package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi, I'm Marc, Hi."

	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("首次出现的位置 %d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the last instance og \"Hi\" is: ")
	fmt.Printf("最后出现的位置 %d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("Replace %s", strings.Replace(str, "Marc", "weishaodaren", 1))
}
