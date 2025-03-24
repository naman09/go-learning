package main

import "fmt"

func main() {
	fmt.Println(add(1, 3))
	fmt.Println(add(3.14, 3))
	fmt.Println(add(1, 2.2))
	fmt.Println(add(-1, 2.2))

}

func add[U, V reals](num1 U, num2 V) float64 {
	return float64(num1) + float64(num2)
}

type reals interface {
	signedInts | unsignedInts | floats
}

type signedInts interface {
	int | int8 | int16 | int32 | int64
}

type unsignedInts interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type floats interface {
	float32 | float64
}
