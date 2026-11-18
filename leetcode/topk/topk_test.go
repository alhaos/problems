package topk

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"example 1", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"single", []int{1}, 1, []int{1}},
		{"all same", []int{4, 4, 4, 4}, 1, []int{4}},
		{"k equals distinct", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"negatives", []int{-1, -1, -2, -2, -2, 3}, 2, []int{-2, -1}},
		{"zero is most frequent", []int{0, 0, 0, 1, 1, 2}, 1, []int{0}},
		{"unordered", []int{3, 0, 1, 0}, 1, []int{0}},
		{"mixed", []int{5, 3, 1, 1, 1, 3, 73, 1}, 2, []int{1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
