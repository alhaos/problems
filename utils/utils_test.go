package utils

import (
	"testing"
)

func TestNextCombination(t *testing.T) {
	tests := []struct {
		name string
		x    uint8
		want uint8
	}{
		{"6 -> 9 (биты 1,2 -> биты 0,3)", 6, 9},
		{"9 -> 10 (биты 0,3 -> биты 1,3)", 9, 10},
		{"10 -> 12 (биты 1,3 -> биты 2,3)", 10, 12},
		{"12 -> 17 (биты 2,3 -> биты 0,4)", 12, 17},
		{"7 -> 11 (биты 0,1,2 -> биты 0,1,3)", 7, 11},
		{"11 -> 13", 11, 13},
		{"13 -> 14", 13, 14},
		{"14 -> 19", 14, 19},
		{"1 -> 2 (один бит)", 1, 2},
		{"2 -> 4", 2, 4},
		{"4 -> 8", 4, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NextCombination(tt.x)
			if got != tt.want {
				t.Errorf("NextCombination(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}
