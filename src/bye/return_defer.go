package main

import "fmt"

// 这可用于在返回语句之后修改返回的 error 时使用
func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func main() {
	fmt.Println(f())
}
