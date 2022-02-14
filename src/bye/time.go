package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main()  {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%04d\n", t.Day(),t.Month(),t.Year())

	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())

	week = 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week)
}