package lis

import (
	"slices"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	ascending := make([]int, 2500)
	for i := range ascending {
		ascending[i] = i
	}

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"example 2", []int{0, 1, 0, 3, 2, 3}, 4},
		{"all equal", []int{7, 7, 7, 7, 7, 7, 7}, 1},
		{"one element", []int{5}, 1},
		{"two ascending", []int{-2, -1}, 2},
		{"two descending", []int{2, 1}, 1},
		{"strictly descending", []int{5, 4, 3, 2, 1}, 1},
		{"strictly ascending", []int{1, 2, 3, 4, 5}, 5},
		{"duplicates inside", []int{4, 10, 4, 3, 8, 9}, 3},
		{"not through the max", []int{10, 22, 9, 33, 21, 50, 41, 60}, 5},
		{"long with repeats", []int{1, 3, 6, 7, 9, 4, 10, 5, 6}, 6},
		{"tail wins", []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}, 6},
		{"max size ascending", ascending, 2500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := slices.Clone(tt.nums)
			got := lengthOfLIS(input)

			if !slices.Equal(input, tt.nums) {
				t.Fatalf("input mutated")
			}

			if got != tt.want {
				t.Fatalf("got %d, want %d", got, tt.want)
			}
		})
	}
}
