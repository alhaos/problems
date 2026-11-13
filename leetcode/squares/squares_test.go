package sortedsquares

import (
	"slices"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example 1", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{"example 2", []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{"single positive", []int{5}, []int{25}},
		{"single negative", []int{-5}, []int{25}},
		{"single zero", []int{0}, []int{0}},
		{"all negative", []int{-5, -3, -1}, []int{1, 9, 25}},
		{"all positive", []int{1, 3, 5}, []int{1, 9, 25}},
		{"all zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"symmetric", []int{-2, -1, 1, 2}, []int{1, 1, 4, 4}},
		{"duplicates", []int{-3, -3, 3, 3}, []int{9, 9, 9, 9}},
		{"max values", []int{-10000, 0, 10000}, []int{0, 100000000, 100000000}},
		{"negatives dominate", []int{-10, -9, -8, 1}, []int{1, 64, 81, 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedSquares(slices.Clone(tt.nums))
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
