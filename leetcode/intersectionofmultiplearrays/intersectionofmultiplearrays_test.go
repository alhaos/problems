package intersectionofmultiplearrays

import (
	"slices"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name string
		nums [][]int
		want []int
	}{
		{"example 1", [][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}}, []int{3, 4}},
		{"example 2", [][]int{{1, 2, 3}, {4, 5, 6}}, []int{}},
		{"single array unsorted", [][]int{{5, 1, 3}}, []int{1, 3, 5}},
		{"single element", [][]int{{7}}, []int{7}},
		{"identical arrays", [][]int{{1, 2, 3}, {3, 2, 1}, {2, 1, 3}}, []int{1, 2, 3}},
		{"one common", [][]int{{1, 2}, {2, 3}, {2, 4}}, []int{2}},
		{"empty intersection in the middle", [][]int{{1, 2}, {3, 4}, {1, 2}}, []int{}},
		{"max value", [][]int{{1000, 1}, {1000, 2}}, []int{1000}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersection(tt.nums)
			if !slices.Equal(got, tt.want) {
				t.Errorf("intersection(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
