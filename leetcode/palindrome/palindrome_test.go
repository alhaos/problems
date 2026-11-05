package palindrome

import (
	"math/rand"
	"strings"
	"testing"
)

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// check проверяет, что got это палиндром, подстрока s и нужной длины.
// Ответов может быть несколько, поэтому сравниваем длину, а не саму строку.
func check(t *testing.T, s, got string, wantLen int) {
	t.Helper()
	if !strings.Contains(s, got) {
		t.Errorf("s=%q: result %q is not a substring", s, got)
		return
	}
	if !isPalindrome(got) {
		t.Errorf("s=%q: result %q is not a palindrome", s, got)
		return
	}
	if len(got) != wantLen {
		t.Errorf("s=%q: result %q has length %d, want %d", s, got, len(got), wantLen)
	}
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		wantLen int
	}{
		{"single", "a", 1},
		{"example1", "babad", 3},
		{"example2", "cbbd", 2},
		{"all same", "aaaa", 4},
		{"two same", "aa", 2},
		{"two different", "ab", 1},
		{"whole string odd", "racecar", 7},
		{"whole string even", "abba", 4},
		{"at the start", "abacdefg", 3},
		{"at the end", "xyzabcba", 5},
		{"in the middle", "xabaxcdefg", 5},
		{"no repeats", "abcdef", 1},
		{"even inside noise", "xyabbaz", 4},
		{"digits and letters", "a1b2b1c", 5},
		{"two candidates same length", "abcbaqxyzyx", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check(t, tt.s, longestPalindrome(tt.s), tt.wantLen)
		})
	}
}

// bruteLen перебирает все подстроки и возвращает длину самого длинного палиндрома.
func bruteLen(s string) int {
	best := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if j-i+1 > best && isPalindrome(s[i:j+1]) {
				best = j - i + 1
			}
		}
	}
	return best
}

func TestLongestPalindromeRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < 20000; i++ {
		b := make([]byte, rng.Intn(14)+1)
		for j := range b {
			b[j] = byte('a' + rng.Intn(3))
		}
		s := string(b)

		check(t, s, longestPalindrome(s), bruteLen(s))
	}
}
