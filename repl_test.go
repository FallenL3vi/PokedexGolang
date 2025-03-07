package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "   hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "BigBananaza  ",
			expected: []string{"bigbananaza"},
		},
		{
			input: "SmallDragon EKWAR  HYUNDAj  ",
			expected: []string{"smalldragon", "ekwar", "hyundaj"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths do not match : %v and %v", actual, c.expected)
			continue
		}
	
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			
			if word != expectedWord {
				t.Errorf("Words do not match : %v and %v", word, expectedWord)
			}
		}
	}

}