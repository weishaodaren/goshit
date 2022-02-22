package main

import "strings"

func main() {
	addBmp := MakeAddSuffix(".bmp")
	println(addBmp("weishaodaren"))
	addJepg := MakeAddSuffix(".jepg")
	println(addJepg("HHHH"))
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
