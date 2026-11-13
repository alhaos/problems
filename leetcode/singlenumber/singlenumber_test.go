package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{2, 2, 1}, 1},
		{"example 2", []int{4, 1, 2, 1, 2}, 4},
		{"single element", []int{1}, 1},
		{"negative", []int{-1, -1, -2}, -2},
		{"zero", []int{0, 5, 5}, 0},
		{"single at the end", []int{3, 7, 3, 9, 7}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.nums); got != tt.want {
				t.Errorf("singleNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
