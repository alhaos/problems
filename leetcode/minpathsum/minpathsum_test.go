package minpathsum

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"example 1", [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{"example 2", [][]int{{1, 2, 3}, {4, 5, 6}}, 12},
		{"single cell", [][]int{{5}}, 5},
		{"single zero", [][]int{{0}}, 0},
		{"one row", [][]int{{1, 2, 3}}, 6},
		{"one column", [][]int{{1}, {2}, {3}}, 6},
		{"2x2", [][]int{{1, 2}, {1, 1}}, 3},
		{"all zeros", [][]int{{0, 0}, {0, 0}}, 0},
		{"down is cheaper", [][]int{{1, 100}, {1, 1}}, 3},
		{"right is cheaper", [][]int{{1, 1}, {100, 1}}, 3},
		{"avoid the middle", [][]int{{1, 1, 1}, {100, 100, 1}, {1, 1, 1}}, 5},
		{"edge column wins", [][]int{{1, 100, 100}, {1, 100, 100}, {1, 1, 1}}, 5},
		{"edge row wins", [][]int{{9, 1, 1, 1}, {9, 9, 9, 1}}, 13},
		{"zigzag", [][]int{{1, 2, 5}, {3, 2, 1}}, 6},
		{"expensive corner", [][]int{{5, 4}, {6, 3}}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.grid); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
