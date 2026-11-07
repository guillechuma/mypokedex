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
			input:    " hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " This IS a StranGe TesT",
			expected: []string{"this", "is", "a", "strange", "test"},
		},
		{
			input:    "     ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("1 expected: %v, got: %v", c.expected, actual)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("2 expected: %v, got: %v", expectedWord, word)
			}
		}
	}
}
