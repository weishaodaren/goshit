package main

import (
	"fmt"
)

func main() {
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	sli = append(sli, 7)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	sli = append(sli, 8, 9, 1, 1, 1, 1, 1, 1)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	sli = append(sli[:3], sli[5:]...)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)
}
