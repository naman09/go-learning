package main

import (
	"fmt"
)

func main() {
	n := 1
	x := 0
	y := 1

	for n < 95 {
		x = x + y
		z := y
		y = x
		x = z
		n++
	}
	fmt.Println("90th fibo value:", x)
}
