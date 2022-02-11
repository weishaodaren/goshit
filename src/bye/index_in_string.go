package main

import (
	"fmt"
	"strings"
)

func main() {
	// string.HasSuffix(suffix, s string) bool
	// string.HasPrefix(suffix, s string) bool
	// strings.Contains(str, s stirng) bool

	var str string = "Hi, I'm Marc, Hi."

	fmt.Printf("The position of \"Marc\" is: ")
	// strings.Index(str, s string) int
	fmt.Printf("首次出现的位置 %d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the last instance og \"Hi\" is: ")
	// strings.LastIndex(str, s string) int
	fmt.Printf("最后出现的位置 %d\n", strings.LastIndex(str, "Hi"))

	// strings.Replace(str, old, new string, n int) string
	fmt.Printf("Replace %s \n", strings.Replace(str, "Marc", "weishaodaren", 1))

	// strings.Count(str, s string) int
	fmt.Printf("这个出现了多少次啊：%d \n", strings.Count(str, "Hi"))

	// strings.Repeat(str, count int) string
	fmt.Printf("The new repeated string is : %s \n", strings.Repeat(str, 100))
}
