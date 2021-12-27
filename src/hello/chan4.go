package main

import (
	"fmt"
	"time"
)

func procuder(ch chan string) {
	fmt.Println("producer start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	fmt.Println("producer end")
}

func customer(ch chan string) {
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func main() {
	fmt.Println("main start")
	ch := make(chan string, 3)
	go procuder(ch)
	go customer(ch)
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
