package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("main start")
	ch := make(chan string, 1)
	ch <- "a" // write

	go func() {
		val := <-ch // read
		fmt.Println(val)
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
