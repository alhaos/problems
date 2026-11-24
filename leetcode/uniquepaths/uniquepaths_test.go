package uniquepaths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m, n int
		want int
	}{
		{"example 1", 3, 7, 28},
		{"example 2", 3, 2, 3},
		{"single cell", 1, 1, 1},
		{"one row", 1, 5, 1},
		{"one column", 5, 1, 1},
		{"2x2", 2, 2, 2},
		{"2x3", 2, 3, 3},
		{"3x3", 3, 3, 6},
		{"3x4", 3, 4, 10},
		{"4x4", 4, 4, 20},
		{"5x5", 5, 5, 70},
		{"10x10", 10, 10, 48620},
		{"max-ish", 23, 12, 193536720},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.m, tt.n); got != tt.want {
				t.Errorf("uniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, got, tt.want)
			}
		})
	}
}
