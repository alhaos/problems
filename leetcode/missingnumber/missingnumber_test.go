package missingnumber

import "testing"

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{3, 0, 1}, 2},
		{"example 2", []int{0, 1}, 2},
		{"example 3", []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{"single, missing 1", []int{0}, 1},
		{"single, missing 0", []int{1}, 0},
		{"missing zero", []int{1, 2, 3}, 0},
		{"missing last", []int{0, 1, 2, 3}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.nums); got != tt.want {
				t.Errorf("missingNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
