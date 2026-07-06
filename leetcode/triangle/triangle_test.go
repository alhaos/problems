package triangle

import "testing"

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		want     int
	}{
		{
			name:     "example from leetcode",
			triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			want:     11,
		},
		{
			name:     "single element",
			triangle: [][]int{{-10}},
			want:     -10,
		},
		{
			name:     "two rows",
			triangle: [][]int{{1}, {2, 3}},
			want:     3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinimumTotal(tt.triangle)
			if got != tt.want {
				t.Errorf("MinimumTotal() = %d, want %d", got, tt.want)
			}
		})
	}
}
