package processor_test

import (
	"go-reloaded/processor"
	"testing"
)

func TestHexToDecimal(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
	}

	tests := []testCase{
		{name: "test valid hex", input: "1E", expected: "30"},
		{name: "test invalid hex", input: "Aqq", expected: "Aqq"},
		{name: "test empty string", input: "", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.input
			processor.HexToDecimal(&s)
			if s != tt.expected {
				t.Errorf("HexToDecimal(%q) = %q; expected -> %q", tt.input, s, tt.expected)
			}
		})
	}
}

func TestBinToDecimal(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
	}

	tests := []testCase{
		{name: "test valid binary", input: "10", expected: "2"},
		{name: "test invalid binary", input: "Aqq", expected: "Aqq"},
		{name: "test empty string", input: "", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.input
			processor.BinToDecimal(&s)
			if s != tt.expected {
				t.Errorf("BinToDecimal(%q) = %q; expected -> %q", tt.input, s, tt.expected)
			}
		})
	}
}
