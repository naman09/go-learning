package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	s := arr[1:4]
	var sCopy = make([]int, 3)
	copy(sCopy, s)
	sCopy[0] = 97
	sCopy[1] = 98
	sCopy[2] = 99

	fmt.Println(arr, len(arr))
	fmt.Println(s, len(s), cap(s))
	fmt.Println(sCopy, len(sCopy), cap(sCopy))

	m := make(map[int]string)
	m[988] = "hnn"
	v, ok := m[99]
	fmt.Println(v, ok)
}
