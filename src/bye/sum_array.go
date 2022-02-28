package main

import "fmt"

func sum(arrF []float32) (sumVal float32) {
	for _, value := range arrF{
		sumVal += value
	}
	return 
}

func sumAndAverage(a int) (int, float32)  {
	sum := 0
	for _, item := range a{
		sum += item
	}
	return sum, float32(sum / len(a))
}

func main()  {
	arrF := []float32{1.0, 2.0, 3.0}
	fmt.Printf("Sum is: %f\n", sum(arrF))
}

