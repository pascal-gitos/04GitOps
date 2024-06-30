package main

import (
	"testing"
)

func TestConcatenate(t *testing.T) {
	result := Concatenate("String ", "concatenate")
	expected := "String concatenate"
	if result != expected {
		t.Error("Concatenation test failed. Expected: %s, got: %s", expected, result)

	}
}

func TestUppercase(t *testing.T) {
	result := Uppercase("string uppercase")
	expected := "STRING UPPERCASE"
	if result != expected {
		printf.Error("Upppercase test failed. Expected: %s, got %s", expected, result)
	}
}
