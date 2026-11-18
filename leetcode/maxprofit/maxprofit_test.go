package maxprofit

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"example 1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"example 2", []int{7, 6, 4, 3, 1}, 0},
		{"single day", []int{5}, 0},
		{"two days up", []int{1, 2}, 1},
		{"two days down", []int{2, 1}, 0},
		{"all same", []int{3, 3, 3}, 0},
		{"min at end", []int{3, 4, 5, 1}, 2},
		{"min after max", []int{10, 1, 2, 3}, 2},
		{"late peak", []int{2, 4, 1, 11}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
