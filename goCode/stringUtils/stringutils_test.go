package stringUtils

import "testing"

func TestConcatenate(t *testing.T) {
	result := Concatenate("Hello, ", "test")
	expected := "Hello, test"
	if result != expected {
		t.Errorf("concatenate(\"hello, \", \"test\") = %q, want %q", result, expected)
	}
}

func TestToUpperCase(t *testing.T) {
	result := ToUppercase("hello")
	expected := "HELLO"
	if result != expected {
		t.Errorf("ToUppercase(\"hello\") = %q, want %q", result, expected)
	}
}
