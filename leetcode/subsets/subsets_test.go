package subsets

import (
	"fmt"
	"slices"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{"one element", []int{0}},
		{"two elements", []int{1, 2}},
		{"three elements", []int{1, 2, 3}},
		{"negatives", []int{-1, 0, 1}},
		{"five elements", []int{5, 4, 3, 2, 1}},
		{"ten elements", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := slices.Clone(tt.nums)
			got := subsets(input)

			if !slices.Equal(input, tt.nums) {
				t.Fatalf("input mutated: %v, was %v", input, tt.nums)
			}

			if want := 1 << len(tt.nums); len(got) != want {
				t.Fatalf("got %d subsets, want %d", len(got), want)
			}

			allowed := make(map[int]bool, len(tt.nums))
			for _, n := range tt.nums {
				allowed[n] = true
			}

			seen := make(map[string]bool, len(got))
			for _, s := range got {
				norm := slices.Clone(s)
				slices.Sort(norm)

				key := fmt.Sprint(norm)
				if seen[key] {
					t.Fatalf("duplicate subset %v", s)
				}
				seen[key] = true

				for i, n := range norm {
					if !allowed[n] {
						t.Fatalf("subset %v has foreign element %d", s, n)
					}
					if i > 0 && norm[i-1] == n {
						t.Fatalf("subset %v has repeated element %d", s, n)
					}
				}
			}
		})
	}
}
