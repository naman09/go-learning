package main

import "fmt"

func main() {
	g := *NewGreeter("Konichiwa")
	//{Konichiwa}
	fmt.Println(g)

	g2 := greeter{greet: "hello"}
	// fmt.Println(g2)

	Print(g2)
	i := 3
	Print(i)
}

func Print(a any) {
	// var a2 fmt.Stringer = a.(fmt.Stringer)
	switch v := a.(type) {
	case fmt.Stringer:
		println(v.String())
	default:
		println("value has no string interface impl")
	}

}

type greeter struct {
	greet string
}

func (g greeter) String() string {
	return "[" + g.greet
}

func NewGreeter(s string) *greeter {
	return &greeter{
		greet: s,
	}
}
