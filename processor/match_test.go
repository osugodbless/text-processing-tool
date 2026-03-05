package processor

import "testing"

func TestMatchPatternCase(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
	}

	tests := []testCase{
		{name: "Multi uppercase pattern", input: "(up, 3)", expected: "(up, 3)"},
		{name: "Multi lowercase pattern", input: "(low, 3)", expected: "(low, 3)"},
		{name: "Multi titlecase pattern", input: "(cap, 3)", expected: "(cap, 3)"},
		{name: "Invalid pattern", input: "cap, 3", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MatchPatternCase(tt.input)
			if s != tt.expected {
				t.Errorf("MatchPatternCase(%q) = %q; expected -> %q", tt.input, s, tt.expected)
			}
		})
	}
}

func TestMatchPatternPunctuation(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
	}

	tests := []testCase{
		{name: "single punctuation match", input: ";", expected: ";"},
		{name: "multi-identical punctuation match", input: "...", expected: "..."},
		{name: "empty string", input: "", expected: ""},
		{name: "Invalid pattern", input: "'", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MatchPatternPunctuation(tt.input)
			if s != tt.expected {
				t.Errorf("MatchPatternPunctuation(%q) = %q; expected -> %q", tt.input, s, tt.expected)
			}
		})
	}
}
