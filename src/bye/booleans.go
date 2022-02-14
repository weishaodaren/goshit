package main

import (
	"fmt"
	"runtime"
)

var prompt = "Enter a digit, e.g. 3" + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main() {
	bool1 := true
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}

	str := ""
	// 判断字符串是否为空
	if str == "" {
		fmt.Println("是否为空啊")
	}

	if str1 := ""; str1 == "" {
		fmt.Println("是否为空啊, 同上")
	}

	if len(str) == 0 {
		fmt.Println("你呢 是否为空啊")
	}

	if runtime.GOOS == "windows" {
		fmt.Println("Windows")
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isGreater(x, y int) bool {
	if x > y {
		return true
	}

	return false
}
