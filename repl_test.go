package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Go is good",
			expected: []string{"go", "is", "good"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Test '%s' failed. Expected: %s", c.input, expectedWord)
				t.Fatalf("Test '%s' failed. Expected: %s", c.input, expectedWord)
			}
		}
	}
}
