package majority

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"single", []int{1}, 1},
		{"example1", []int{3, 2, 3}, 3},
		{"example2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"all same", []int{7, 7, 7, 7}, 7},
		{"negatives", []int{-1, -1, 2, -1, 3}, -1},
		{"majority at end", []int{1, 2, 3, 3, 3}, 3},
		{"two elements", []int{5, 5}, 5},
		{"minimal majority", []int{1, 2, 1, 2, 1}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.nums); got != tt.want {
				t.Errorf("majorityElement(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
