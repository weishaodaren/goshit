package main

import (
	"fmt"
	"time"
)

func main() {
	GoA()
	time.Sleep(1 * time.Second)
	fmt.Println("main")
}

func GoA() {
	defer (func() {
		if err := recover(); err != nil {
			fmt.Println("panuc:" + fmt.Sprintf("%s", err))
		}
	})()

	go GoB()
}

func GoB() {
	panic("error")

}
