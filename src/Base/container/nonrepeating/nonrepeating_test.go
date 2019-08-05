package main

import (
	"testing"
)

func TestSubstr(t *testing.T)  {
	tests := []struct{
		s string
		ans int
	} {
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},
		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese cases
		{"黑化肥会发会发挥", 5},
	}

	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d ", actual, tt.s, tt.ans)
		}
	}
}