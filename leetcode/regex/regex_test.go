package regex

import (
	"math/rand"
	"regexp"
	"strings"
	"testing"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name string
		s, p string
		want bool
	}{
		// без спецсимволов
		{"exact single", "a", "a", true},
		{"different letter", "a", "b", false},
		{"same word", "abc", "abc", true},
		{"different word", "abc", "abd", false},
		{"pattern shorter", "abc", "ab", false},
		{"pattern longer", "ab", "abc", false},

		// точка
		{"dot single", "a", ".", true},
		{"dot in middle", "abc", "a.c", true},
		{"dot mismatch", "abc", "a.d", false},
		{"only dots", "abc", "...", true},
		{"too many dots", "ab", "...", false},

		// звёздочка
		{"star zero", "a", "ab*", true},
		{"star one", "ab", "ab*", true},
		{"star many", "abbb", "ab*", true},
		{"example1", "aa", "a", false},
		{"example2", "aa", "a*", true},
		{"star must cover all", "ab", "a*", false},
		{"star then same letter", "aaa", "a*a", true},
		{"star needs one more", "a", "ab*a", false},

		// точка со звёздочкой
		{"example3", "ab", ".*", true},
		{"dot star in middle", "abcd", "a.*d", true},
		{"dot star zero", "ad", "a.*d", true},
		{"dot star too short", "ab", ".*c", false},

		// сложные
		{"example4", "aab", "c*a*b", true},
		{"mississippi no", "mississippi", "mis*is*p*.", false},
		{"mississippi yes", "mississippi", "mis*is*ip*.", true},
		{"several stars", "aaa", "ab*a*c*a", true},
		{"dot star then star", "bbbba", ".*a*a", true},
		{"dot star not enough", "a", ".*..a*", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("isMatch(%q, %q) = %v, want %v", tt.s, tt.p, got, tt.want)
			}
		})
	}
}

// randomPattern собирает валидный шаблон: символ из {a, b, .}
// и необязательная звёздочка после него.
func randomPattern(rng *rand.Rand) string {
	letters := []string{"a", "b", "."}
	var sb strings.Builder
	for i, n := 0, rng.Intn(6)+1; i < n; i++ {
		sb.WriteString(letters[rng.Intn(len(letters))])
		if rng.Intn(3) == 0 {
			sb.WriteByte('*')
		}
	}
	return sb.String()
}

func randomString(rng *rand.Rand) string {
	b := make([]byte, rng.Intn(8)+1)
	for i := range b {
		b[i] = byte('a' + rng.Intn(2))
	}
	return string(b)
}

func TestIsMatchRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < 50000; i++ {
		s, p := randomString(rng), randomPattern(rng)

		want := regexp.MustCompile("^(?:" + p + ")$").MatchString(s)
		if got := isMatch(s, p); got != want {
			t.Fatalf("isMatch(%q, %q) = %v, want %v", s, p, got, want)
		}
	}
}
