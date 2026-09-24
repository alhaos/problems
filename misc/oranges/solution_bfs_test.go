package oranges

import "testing"

func TestOrangesRottingBFS(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"example 1", [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{"unreachable orange", [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
		{"no fresh, has rotten", [][]int{{0, 2}}, 0},
		{"only empty", [][]int{{0}}, 0},
		{"only fresh", [][]int{{1}}, -1},
		{"only rotten", [][]int{{2}}, 0},
		{"all rotten", [][]int{{2, 2}, {2, 2}}, 0},
		{"blocked by empty cell", [][]int{{2, 0, 1}}, -1},
		{"single row", [][]int{{2, 1, 1, 1, 1}}, 4},
		{"two sources rot in parallel", [][]int{{2, 1, 1}, {1, 1, 1}, {1, 1, 2}}, 2},
		{"sources rot simultaneously", [][]int{{2, 2}, {1, 1}, {0, 0}, {2, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrangesRottingBFS(clone(tt.grid)); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestOrangesRottingBFSLarge(t *testing.T) {
	g := make([][]int, 10)
	for i := range g {
		g[i] = make([]int, 10)
		for j := range g[i] {
			g[i][j] = 1
		}
	}
	g[0][0] = 2
	if got := OrangesRottingBFS(g); got != 18 {
		t.Errorf("got %d, want 18", got)
	}
}
