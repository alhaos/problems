package maxsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"single positive", []int{1}, 1},
		{"all positive", []int{5, 4, -1, 7, 8}, 23},
		{"all negative", []int{-3, -1, -2}, -1},
		{"single negative", []int{-1}, -1},
		{"single zero", []int{0}, 0},
		{"two negatives", []int{-2, -1}, -1},
		{"negative in middle", []int{2, -1, 2}, 3},
		{"alternating", []int{1, -1, 1, -1, 1}, 1},
		{"restart after dip", []int{-1, 2, -1, 2}, 3},
		{"big dip not worth", []int{3, -10, 4}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
