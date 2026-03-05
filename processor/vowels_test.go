package processor_test

import (
	"go-reloaded/processor"
	"testing"
)

func TestIsVowel(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected bool
	}

	tests := []testCase{
		{name: "test valid string", input: "Incubator", expected: true},
		{name: "test invalid string", input: "zainab", expected: false},
		{name: "test empty string", input: "", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := processor.IsVowel(&tt.input)
			if s != tt.expected {
				t.Errorf("IsVowel(%q) returned a value not expected", tt.input)
			}
		})
	}
}
