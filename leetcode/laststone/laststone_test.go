package laststone

import "testing"

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{"example 1", []int{2, 7, 4, 1, 8, 1}, 1},
		{"single", []int{1}, 1},
		{"two equal", []int{2, 2}, 0},
		{"two different", []int{1, 3}, 2},
		{"three", []int{3, 7, 2}, 2},
		{"all destroyed", []int{9, 3, 2, 10}, 0},
		{"four equal", []int{2, 2, 2, 2}, 0},
		{"one big", []int{100, 1, 1, 1}, 97},
		{"two big equal", []int{10, 4, 2, 10}, 2},
		{"ones and two", []int{1, 1, 2}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make([]int, len(tt.stones))
			copy(in, tt.stones)
			if got := lastStoneWeight(in); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
