package acmp45

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{10, 25},
		{13, -1},
		{8, 8},
		{90, 259},
		{0, 10},
		{1, 1},
		{7, 7},
		{11, -1},
		{100, 455},
		{1_000_000_000, 555555555888},
	}
	for _, tt := range tests {
		if got := Solve(tt.n); got != tt.want {
			t.Errorf("Solve(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
