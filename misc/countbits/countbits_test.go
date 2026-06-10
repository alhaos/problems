package countbits

import "testing"

func TestCountBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"two", 2, 1},   // 10
		{"three", 3, 2}, // 11
		{"seven", 7, 3}, // 111
		{"eight", 8, 1}, // 1000
		{"thirteen", 13, 3},  // 1101
		{"255", 255, 8},      // 11111111
		{"256", 256, 1},      // 100000000
		{"large", 1<<20 - 1, 20}, // 20 ones
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountBits(tc.n)
			if got != tc.want {
				t.Errorf("CountBits(%d) = %d, want %d", tc.n, got, tc.want)
			}
		})
	}
}
