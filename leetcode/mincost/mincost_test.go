package mincost

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{"example 1", []int{10, 15, 20}, 15},
		{"example 2", []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
		{"two zeros", []int{0, 0}, 0},
		{"two equal", []int{5, 5}, 5},
		{"two, first cheaper", []int{1, 2}, 1},
		{"two, second cheaper", []int{2, 1}, 1},
		{"four", []int{10, 15, 20, 25}, 30},
		{"all ones", []int{1, 1, 1, 1}, 2},
		{"zero at ends", []int{0, 1, 2, 0}, 1},
		{"expensive ends", []int{100, 1, 1, 100}, 2},
		{"free path", []int{0, 0, 0, 1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.cost); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
