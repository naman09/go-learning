package main

import "fmt"

func main() {
	const (
		OtherExecute = 1 << iota
		OtherWrite
		OtherRead
		GroupExecute
		GroupWrite
		GroupRead
		OwnerExecute
		OwnerWrite
		OwnerRead
	)

	fmt.Printf("%O\n", OwnerExecute)
	fmt.Printf("%O", OtherExecute+OtherRead+GroupExecute+GroupRead+OwnerExecute+OwnerRead)
}
