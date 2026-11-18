package topkfrequentelements

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int // в отсортированном виде
	}{
		{"example 1", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"example 2", []int{1}, 1, []int{1}},
		{"k = 1", []int{4, 4, 4, 5, 5, 6}, 1, []int{4}},
		{"k = all unique", []int{1, 2, 3, 1, 2, 1}, 3, []int{1, 2, 3}},
		{"negatives", []int{-1, -1, -2, -2, -2, 3}, 2, []int{-2, -1}},
		{"all same", []int{7, 7, 7, 7}, 1, []int{7}},
		{"bounds", []int{-10000, 10000, 10000, -10000, 10000}, 1, []int{10000}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(slices.Clone(tt.nums), tt.k)
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("topKFrequent(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
