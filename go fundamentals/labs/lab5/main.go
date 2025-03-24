package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toUpper("acc", "bbB"))
	fmt.Println(toUpper())
	inp := []string{"x", "y"}
	fmt.Println(toUpper(inp...), inp)
}

func toUpper(lstStrings ...string) []string {
	var result []string
	for _, currentString := range lstStrings {
		result = append(result, strings.ToUpper(currentString))
	}
	return result
}
