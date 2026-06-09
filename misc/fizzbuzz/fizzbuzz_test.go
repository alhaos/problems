package fizzbuzz

import "testing"

func TestFizzBuzzSum(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{5, 7},
		{10, 22},
		{15, 60},
		{0, 0},
	}

	for _, tt := range tests {
		got := FizzBuzzSum(tt.n)
		if got != tt.want {
			t.Errorf("FizzBuzzSum(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
