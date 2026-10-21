package stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"n=1", 1, 1},
		{"n=2", 2, 2},
		{"n=3", 3, 3},
		{"n=4", 4, 5},
		{"n=5", 5, 8},
		{"n=10", 10, 89},
		{"n=20", 20, 10946},
		{"n=45 max", 45, 1836311903},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func referenceClimbStairs(n int) int {
	dp := make([]int, n+2)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func TestClimbStairsAll(t *testing.T) {
	for n := 1; n <= 45; n++ {
		if got, want := climbStairs(n), referenceClimbStairs(n); got != want {
			t.Fatalf("n=%d: got %d, want %d", n, got, want)
		}
	}
}
