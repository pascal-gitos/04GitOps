package main

import (
	"fmt"
	"goCode/stringUtils"
)

func main() {
	str1 := "hello"
	str2 := "goUtils"
	concatenated := stringUtils.Concatenate(str1, str2)
	fmt.Println("Concatenate string:", concatenated)

	lowercase := "lower string"
	uppercase := stringUtils.ToUppercase(lowercase)
	fmt.Println("uppercase string:", uppercase)
}
