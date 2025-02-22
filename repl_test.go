package main

import (
	"testing"
)

// All tests go inside TestXXX functions that take a *testing.T argument

func TestCleanInput(t *testing.T) {
	// creating a slice a test case structs
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello WORld",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hElLo WORld  ",
			expected: []string{"hello", "world"},
		},
	}

	// loop over the cases and run the tests
	for _, c := range cases {
		actual := cleanInput(c.input)
		// checking the length of the actual slice
		if len(actual) != len(c.expected) {
			t.Errorf("Length of returned from cleanInput: %v doesnt match length expected: %v", len(actual), len(c.expected))
			t.Fail()
		}
		// checking each word in the slice
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Actual word: %v doesnt match expected word: %v", word, expectedWord)
				t.Fail()
			}
		}
	}
}
