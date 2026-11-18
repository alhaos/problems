package maxproduct

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{2, 3, -2, 4}, 6},
		{"example 2", []int{-2, 0, -1}, 0},
		{"single negative", []int{-2}, -2},
		{"single zero", []int{0}, 0},
		{"two zeros", []int{0, 0}, 0},
		{"zero then positive", []int{0, 2}, 2},
		{"two negatives", []int{-1, -1}, 1},
		{"neg pos neg", []int{-2, 3, -4}, 24},
		{"three negatives", []int{-1, -2, -3}, 6},
		{"negative in middle", []int{3, -1, 4}, 4},
		{"two negatives late", []int{2, -5, -2, -4, 3}, 24},
		{"zero splits", []int{-2, -3, 0, 2, -5}, 6},
		{"zero before pair", []int{2, 3, 0, -2, -4}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.nums); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
