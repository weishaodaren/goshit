package main

import "fmt"

func main()  {
	function1()
}

func function1() {
	fmt.Println("In function1 at the top")
	defer f()
	defer a()
	defer function2()
	fmt.Println("In function1 at the bottom")
}

func function2()  {
	fmt.Println("Function2: Deferred util the end of the calling function")
}

func a()  {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func f()  {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}