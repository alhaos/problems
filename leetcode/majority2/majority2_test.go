package majority2

import (
	"slices"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"single", []int{1}, []int{1}},
		{"example1", []int{3, 2, 3}, []int{3}},
		{"two elements", []int{1, 2}, []int{1, 2}},
		{"all same", []int{2, 2, 2}, []int{2}},
		{"two winners", []int{1, 1, 1, 3, 3, 2, 2, 2}, []int{1, 2}},
		{"two winners 2", []int{2, 2, 1, 1, 1, 2, 2}, []int{1, 2}},
		{"negatives", []int{-1, -1, -1, 2, 2, 2, 3}, []int{-1, 2}},
		{"one winner among noise", []int{2, 2, 1, 3}, []int{2}},
		{"zero is not a winner", []int{0, 3, 3, 1, 2}, []int{3}},
		{"none, all distinct", []int{1, 2, 3}, []int{}},
		{"none, candidates need verification", []int{1, 2, 3, 4, 1, 2, 3, 4}, []int{}},
		{"none, tail noise", []int{1, 1, 2, 3, 4, 5, 6}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := slices.Clone(tt.nums)
			got := majorityElement(sc)
			slices.Sort(got)
			want := slices.Clone(tt.want)
			slices.Sort(want)

			if len(got) != len(want) || (len(got) > 0 && !slices.Equal(got, want)) {
				t.Errorf("majorityElement(%v) = %v, want %v", tt.nums, got, want)
			}
		})
	}
}
