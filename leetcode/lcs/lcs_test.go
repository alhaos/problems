package lcs

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name         string
		text1, text2 string
		want         int
	}{
		{"example 1", "abcde", "ace", 3},
		{"example 2", "abc", "abc", 3},
		{"example 3", "abc", "def", 0},
		{"single equal", "a", "a", 1},
		{"single different", "a", "b", 0},
		{"reversed", "abc", "cba", 1},
		{"pair reversed", "AB", "BA", 1},
		{"one inside another", "abc", "ahbgdc", 3},
		{"second is subsequence", "abcd", "bd", 2},
		{"repeats", "aaaa", "aa", 2},
		{"gap in the middle", "aab", "azb", 2},
		{"classic 1", "abcdgh", "aedfh", 3},
		{"classic 2", "AGGTAB", "GXTXAYB", 4},
		{"horse and ros", "horse", "ros", 2},
		{"one common letter", "bl", "yby", 1},
		{"palindrome inside", "abcba", "abcbcba", 5},
		{"scattered", "oxcpqrsvwf", "shmtulqrypy", 2},
		{"leetcode tricky", "ezupkr", "ubmrapg", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.text1, tt.text2); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
