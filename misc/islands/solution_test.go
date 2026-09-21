package islands

import "testing"

func toGrid(rows ...string) [][]byte {
	g := make([][]byte, len(rows))
	for i, r := range rows {
		g[i] = []byte(r)
	}
	return g
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{"one big island", toGrid("11110", "11010", "11000", "00000"), 1},
		{"three islands", toGrid("11000", "11000", "00100", "00011"), 3},
		{"all water", toGrid("000", "000"), 0},
		{"all land", toGrid("111", "111"), 1},
		{"single cell land", toGrid("1"), 1},
		{"single cell water", toGrid("0"), 0},
		{"diagonal is not connected", toGrid("10", "01"), 2},
		{"snake", toGrid("10111", "10101", "11101"), 1},
		{"checkerboard", toGrid("1010", "0101", "1010"), 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumIslands(tt.grid); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestNumIslandsLarge(t *testing.T) {
	// 300x300 суши — проверка на глубину рекурсии / производительность
	g := make([][]byte, 300)
	for i := range g {
		g[i] = make([]byte, 300)
		for j := range g[i] {
			g[i][j] = '1'
		}
	}
	if got := NumIslands(g); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}
