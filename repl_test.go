package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"Hello, World!", []string{"Hello", "World"}},
		{"", []string{}},
		{"   ", []string{}},
		{"Hello, World! How are you?", []string{"Hello", "World", "How", "are", "you"}},
	}

	for _, tc := range cases {
		actual := CleanInput(tc.input)
		if len(actual) != len(tc.expected) {
			t.Errorf("cleanInput(%q) = %v, want %v", tc.input, actual, tc.expected)
		}
	}
}
