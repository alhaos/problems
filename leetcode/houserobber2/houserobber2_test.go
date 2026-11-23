package houserobber2

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{2, 3, 2}, 3},
		{"example 2", []int{1, 2, 3, 1}, 4},
		{"example 3", []int{1, 2, 3}, 3},
		{"single", []int{5}, 5},
		{"single zero", []int{0}, 0},
		{"two", []int{1, 2}, 2},
		{"two reversed", []int{2, 1}, 2},
		{"all zeros", []int{0, 0, 0}, 0},
		{"differs from linear", []int{2, 7, 9, 3, 1}, 11},
		{"both ends big", []int{5, 1, 1, 5}, 6},
		{"big at the end", []int{1, 3, 1, 3, 100}, 103},
		{"mixed", []int{4, 1, 2, 7, 5, 3, 1}, 14},
		{"four ones and two", []int{1, 1, 1, 2}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.nums); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
