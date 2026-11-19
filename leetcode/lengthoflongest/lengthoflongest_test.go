package lengthoflongest

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"example 1", "abcabcbb", 3},
		{"example 2", "bbbbb", 1},
		{"example 3", "pwwkew", 3},
		{"empty", "", 0},
		{"single space", " ", 1},
		{"single char", "a", 1},
		{"two distinct", "au", 2},
		{"all distinct", "abcdef", 6},
		{"repeat jumps back", "dvdf", 3},
		{"stale index", "abba", 2},
		{"tricky", "tmmzuxt", 5},
		{"with spaces", "a b c a", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
