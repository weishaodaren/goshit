package main

import (
	"fmt"
	"os"
)

func main() {
	var month int = 13
	seon, err := season(month)
	if !err {
		fmt.Println("啊啊啊啊啊啊")
		os.Exit(1)
	}
	fmt.Printf("Season is %s \n",seon)

}

func season(m int) (string, bool) {
	switch m {
	case 12, 1, 2:
		return "Winter" ,true

	case 3, 4, 5:
		return "Spring" ,true

	case 6, 7, 8:
		return "Summer" ,true

	case 9, 10, 11:
		return "Autumn" ,true

	default :
		return "NNNNNN" ,false

	}

	return "Season unknown" ,true

}
