package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("main start")
	ch := make(chan string)
	go func() {
		ch <- "a"
	}()
	go func() {
		val := <-ch
		fmt.Println(val)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
