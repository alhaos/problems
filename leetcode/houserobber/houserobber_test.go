package houserobber

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// {"example 1", []int{1, 2, 3, 1}, 4},
		{"example 2", []int{2, 7, 9, 3, 1}, 12},
		{"single", []int{5}, 5},
		{"single zero", []int{0}, 0},
		{"two, first bigger", []int{2, 1}, 2},
		{"two, second bigger", []int{1, 2}, 2},
		{"all zeros", []int{0, 0, 0}, 0},
		{"edges win", []int{2, 1, 1, 2}, 4},
		{"skip two in a row", []int{3, 10, 3, 1, 2}, 12},
		{"big at both ends", []int{100, 1, 1, 100}, 200},
		{"big at the end", []int{1, 3, 1, 3, 100}, 103},
		{"mixed", []int{4, 1, 2, 7, 5, 3, 1}, 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.nums); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
