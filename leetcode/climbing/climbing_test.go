package climbing

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"one step", 1, 1},
		{"two steps", 2, 2},
		{"three steps", 3, 3},
		{"four steps", 4, 5},
		{"five steps", 5, 8},
		{"seven steps", 7, 21},
		{"ten steps", 10, 89},
		{"twenty steps", 20, 10946},
		{"thirty steps", 30, 1346269},
		{"max n", 45, 1836311903},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
