package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("main start")
	ch := make(chan string, 3)
	go producer(ch)

	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

func producer(ch chan string) {
	fmt.Println("procuder start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	fmt.Println("producer end")
}
