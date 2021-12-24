package main

import "fmt"

func main() {
	fmt.Println("beigin")

	for i := 1; i <= 10; i++ {
		if i == 6 {
			// break
			// continue
			goto END
		}

		fmt.Println("i =", i)
	}

END:
	fmt.Println("end")
}
