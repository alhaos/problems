package movezeroes

import (
	"slices"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"example", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"single zero", []int{0}, []int{0}},
		{"single nonzero", []int{5}, []int{5}},
		{"no zeros", []int{1, 2, 3}, []int{1, 2, 3}},
		{"all zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"zeros at end", []int{1, 2, 0, 0}, []int{1, 2, 0, 0}},
		{"zeros at start", []int{0, 0, 1, 2}, []int{1, 2, 0, 0}},
		{"negatives", []int{0, -1, 0, -2, 3}, []int{-1, -2, 3, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.in)
			if !slices.Equal(tt.in, tt.want) {
				t.Errorf("got %v, want %v", tt.in, tt.want)
			}
		})
	}
}
