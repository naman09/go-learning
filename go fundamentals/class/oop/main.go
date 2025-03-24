package main

import "fmt"

func main() {
	s := user{
		id:   "1",
		name: "naman",
	}
	fmt.Println(s)
}

type user struct {
	id   string
	name string
}
