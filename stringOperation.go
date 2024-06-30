package main

import (
	"fmt"
	"strings"
)

func Concatenate(a, b string) string {
	return a + b
}

func Uppercase(input string) string {
	return strings.ToUpper(input)
}

func main1() {
	fmt.Println("Concatenated strings:", Concatenate("String", "concatenate"))
	fmt.Println("Uppercase string:", Uppercase("string uppercase"))
}
