package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{

		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " this is a test ",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "goodbye-world",
			expected: []string{"goodbye-world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of %v doent match length of %v", actual, c.expected)
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Word %v dont match with %v", actual[i], c.expected[i])
			}
		}
	}
}
