package main

import (
	"testing"
)

func TestClearInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := clearInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("the lenghts are not equal: %v vs %v",
				len(actual),
				len(cs.expected),
			)
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("words are not equal: %v vs %v",
					actualWord,
					expectedWord,
				)
			}
		}
	}
}
