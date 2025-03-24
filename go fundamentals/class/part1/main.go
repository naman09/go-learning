package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	// fmt.Println("Hello world")
	var myInt int
	var foo = "0yyy\n\r"
	ch := "a"
	fmt.Println(foo, myInt, ch)

	myPtr := &myInt
	fmt.Println(*myPtr)

	interpertedString := "apple\npeach\nplum"
	rawString := `apple\npeach\nplum`

	jsonString := `{
	"username": "naman09"
	"age": 21
}`

	fmt.Println(interpertedString, rawString)
	fmt.Println(jsonString)

	//int, uint, float32, float64, complex64, complex128
	floatVar := 22.2
	intVar := int(floatVar)
	fmt.Println(intVar)

	const a = 3
	const b = a * 40
	fmt.Println(b)

	const (
		OtherExecute = 1 << iota
		OtherWrite
		OtherRead
	)
	fmt.Println("%O\n", OtherWrite+OtherExecute+OtherRead)

	/* Aggregate Types */

	/* 1. Arrays */
	arr := [3]int{1, 2, 3} //size is fixed
	arr2 := arr            //copied by value
	arr2[0] = 34
	fmt.Println(arr, arr2)
	fmt.Println(arr == arr2) //comparable

	//Slices
	var arr3 []int
	fmt.Printf("%T\n", arr3)

	fmt.Println(append(arr3, 1, 43))

	s2 := []int{1, 2, 3}
	// slices.Delete(s2, 1, 3)
	fmt.Println(slices.Delete(s2, 2, 3))

	//struct

	type myStruct struct {
		id   int
		name string
	}
	var s struct {
		id   int
		name string
		st3  myStruct
	}
	fmt.Printf("%#v\n", s)
	st2 := myStruct{id: 42, name: "add"}
	fmt.Println(st2)

	if i := 2; i < 0 { //used with map: if _, ok = m["key"]; ok {}
		fmt.Println("dd")
	} else {
		fmt.Println("else block")
	}

}
