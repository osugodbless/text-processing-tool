package processor_test

import (
	"go-reloaded/processor"
	"testing"
)

func TestExtractDigit(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected int
	}

	tests := []testCase{
		{name: "Single digit extraction", input: "cap 1", expected: 1},
		{name: "multiple digit extraction", input: "ca2p 1", expected: 21},
		{name: "empty string", input: "", expected: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := processor.ExtractDigit(tt.input)
			if s != tt.expected {
				t.Errorf("ExtractDigit(%q) = %q; expected -> %q", tt.input, s, tt.expected)
			}
		})
	}
}
