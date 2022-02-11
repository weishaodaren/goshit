package main

import (
	"fmt"
	"math"
)

func main() {
	var n int16 = 34
	var m int32

	m = int32(n)

	fmt.Println("32 bit int is %d\n", m)
	fmt.Println("16 bit int is %d\n", n)

	print(Uint8FromInt(123321))
	fmt.Println("asd \n",IntFromFloat64(1.12321312))
}

func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the unit8 range", n)
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= .5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}
