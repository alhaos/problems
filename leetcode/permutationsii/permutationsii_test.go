package permutationsii

import (
	"fmt"
	"slices"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int // ожидаемое число уникальных перестановок
	}{
		{"one element", []int{1}, 1},
		{"all same", []int{2, 2, 2}, 1},
		{"example 1", []int{1, 1, 2}, 3},
		{"no duplicates", []int{1, 2, 3}, 6},
		{"unsorted with dups", []int{2, 1, 2, 1}, 6},
		{"three kinds", []int{1, 1, 2, 2, 3}, 30},
		{"negatives", []int{-1, 2, -1, 2}, 6},
		{"eight with dups", []int{1, 1, 2, 2, 3, 3, 4, 4}, 2520},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := slices.Clone(tt.nums)
			got := permuteUnique(input)

			if !slices.Equal(input, tt.nums) {
				t.Fatalf("input mutated: %v, was %v", input, tt.nums)
			}

			if len(got) != tt.want {
				t.Fatalf("got %d permutations, want %d", len(got), tt.want)
			}

			sortedWant := slices.Clone(tt.nums)
			slices.Sort(sortedWant)

			seen := make(map[string]bool, len(got))
			for _, p := range got {
				key := fmt.Sprint(p)
				if seen[key] {
					t.Fatalf("duplicate permutation %v", p)
				}
				seen[key] = true

				sorted := slices.Clone(p)
				slices.Sort(sorted)
				if !slices.Equal(sorted, sortedWant) {
					t.Fatalf("permutation %v has wrong elements", p)
				}
			}
		})
	}
}
