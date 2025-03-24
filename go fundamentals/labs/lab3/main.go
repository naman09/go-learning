package main

import "fmt"

func main() {
	m := map[string]bool{
		"adent":      true,
		"tmacmillan": true,
	}

	fmt.Println(m)
	var _, ok = m["adent"]
	fmt.Println("adent present: ", ok)

	fmt.Println("fprefect present:", mapHas(m, "fprefect"))

	m["fpresent"] = true
	fmt.Println(m)

}
func mapHas(m map[string]bool, key string) bool {
	_, ok := m[key]
	return ok
}
