package uniquepaths2

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"example 1", [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{"example 2", [][]int{{0, 1}, {0, 0}}, 1},
		{"single free", [][]int{{0}}, 1},
		{"single blocked", [][]int{{1}}, 0},
		{"one row free", [][]int{{0, 0, 0}}, 1},
		{"one column free", [][]int{{0}, {0}, {0}}, 1},
		{"one row blocked", [][]int{{0, 1, 0}}, 0},
		{"start blocked", [][]int{{1, 0}, {0, 0}}, 0},
		{"end blocked", [][]int{{0, 0}, {0, 1}}, 0},
		{"both diagonals blocked", [][]int{{0, 1}, {1, 0}}, 0},
		{"no obstacles 3x3", [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 6},
		{"obstacle in first row", [][]int{{0, 0, 1, 0}, {0, 0, 0, 0}}, 2},
		{"obstacle in first column", [][]int{{0, 0}, {1, 0}, {0, 0}}, 1},
		{"wall with a gap", [][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}}, 1},
		{"full wall", [][]int{{0, 0}, {1, 1}, {0, 0}}, 0},
		{"one obstacle in 3x4", [][]int{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}}, 4},
		{"obstacle before the end", [][]int{{0, 0, 0}, {0, 0, 0}, {0, 1, 0}}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.grid); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
