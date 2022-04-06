package main

import (
	"fmt"
	"match"
	"match/big"
)

func main() {
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1965)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Prinf("Big Int: %v \n", ip)
}
