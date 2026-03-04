package processor_test

import (
	"go-reloaded/processor"
	"testing"
)

func TestLowercase(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
	}

	tests := []testCase{
		{name: "test uppercase string", input: "INPUT", expected: "input"},
		{name: "test mixedcase string", input: "InPuT", expected: "input"},
		{name: "test titlecase string", input: "Input", expected: "input"},
		{name: "test string containing digits", input: "12345", expected: "12345"},
		{name: "test empty string", input: "", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.input
			processor.Lowercase(&s)
			if s != tt.expected {
				t.Errorf("Lowercase(%q) = %q; expected -> %q", tt.input, s, tt.expected)
			}
		})
	}
}

func TestLowercaseMutatesCorrectly(t *testing.T) {
	s := "INPUT"
	processor.Lowercase(&s)
	if s != "input" {
		t.Errorf("Lowercase did not mutate correctly")
	}
}

func TestUppercase(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
	}

	tests := []testCase{
		{name: "test Lowercase string", input: "great", expected: "GREAT"},
		{name: "test mixedcase string", input: "BolaJI", expected: "BOLAJI"},
		{name: "test single character string", input: "i", expected: "I"},
		{name: "test empty string", input: "", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.input
			processor.Uppercase(&s)
			if s != tt.expected {
				t.Errorf("Uppercase(%q) = %q; expected -> %q", tt.input, s, tt.expected)
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
	}

	tests := []testCase{
		{name: "test single character string", input: "i", expected: "I"},
		{name: "test lowercase string", input: "zainab", expected: "Zainab"},
		{name: "test empty string", input: "", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.input
			processor.Capitalize(&s)
			if s != tt.expected {
				t.Errorf("Capitalize(%q) = %q; expected -> %q", tt.input, s, tt.expected)
			}
		})
	}
}
