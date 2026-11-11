package combinationsum

import (
	"fmt"
	"slices"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{"example 1", []int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{"example 2", []int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{"no solution", []int{2}, 1, [][]int{}},
		{"single candidate fits", []int{7}, 7, [][]int{{7}}},
		{"single candidate repeated", []int{3}, 9, [][]int{{3, 3, 3}}},
		{"unsorted input", []int{7, 3, 2, 6}, 7, [][]int{{2, 2, 3}, {7}}},
		{"many combinations", []int{2, 3, 4}, 9, [][]int{
			{2, 2, 2, 3}, {2, 3, 4}, {3, 3, 3},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := slices.Clone(tt.candidates)
			got := combinationSum(input, tt.target)

			if !slices.Equal(input, tt.candidates) {
				t.Fatalf("input mutated: %v, was %v", input, tt.candidates)
			}

			if !equalCombinations(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// equalCombinations сравнивает наборы комбинаций без учёта порядка
// комбинаций и порядка чисел внутри каждой.
func equalCombinations(got, want [][]int) bool {
	return fmt.Sprint(normalize(got)) == fmt.Sprint(normalize(want))
}

func normalize(c [][]int) []string {
	keys := make([]string, 0, len(c))
	for _, comb := range c {
		s := slices.Clone(comb)
		slices.Sort(s)
		keys = append(keys, fmt.Sprint(s))
	}
	slices.Sort(keys)
	return keys
}
