package main

import (
	"fmt"
)

const c = "C"

var v int = 5

type T struct{}

func init() {}

func main() {
	var a int
	Func1()

	fmt.Println(a)
}

func (t T) Method1() {

}

// 只有当某个函数需要被外部包调用的时候才使用大写字母开头
func Func1() {}
