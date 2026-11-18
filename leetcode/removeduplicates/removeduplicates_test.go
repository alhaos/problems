package removeduplicates

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"example", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
		{"short", []int{1, 1, 2}, []int{1, 2}},
		{"single", []int{7}, []int{7}},
		{"all same", []int{2, 2, 2, 2}, []int{2}},
		{"no duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"negatives", []int{-3, -3, -1, 0, 0, 5}, []int{-3, -1, 0, 5}},
		{"two elements same", []int{4, 4}, []int{4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := removeDuplicates(tt.in)
			if k != len(tt.want) {
				t.Fatalf("k = %d, want %d", k, len(tt.want))
			}
			if !slices.Equal(tt.in[:k], tt.want) {
				t.Errorf("nums[:k] = %v, want %v", tt.in[:k], tt.want)
			}
		})
	}
}
