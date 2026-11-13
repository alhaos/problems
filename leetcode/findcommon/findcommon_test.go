package findcommon

import (
	"slices"
	"testing"
)

func TestCommonChars(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{"example 1", []string{"bella", "label", "roller"}, []string{"e", "l", "l"}},
		{"example 2", []string{"cool", "lock", "cook"}, []string{"c", "o"}},
		{"single word", []string{"abc"}, []string{"a", "b", "c"}},
		{"single word with repeats", []string{"aabbc"}, []string{"a", "a", "b", "b", "c"}},
		{"no common", []string{"abc", "def"}, []string{}},
		{"identical", []string{"abc", "abc", "abc"}, []string{"a", "b", "c"}},
		{"min of counts", []string{"aaa", "aa", "aaaa"}, []string{"a", "a"}},
		{"one letter missing in one word", []string{"abc", "abd", "xyz"}, []string{}},
		{"all same letter", []string{"zzzz", "zzz"}, []string{"z", "z", "z"}},
		{"alphabet edges", []string{"az", "za", "aazz"}, []string{"a", "z"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := commonChars(slices.Clone(tt.words))
			slices.Sort(got)
			slices.Sort(tt.want)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
