package coinchange

import (
	"slices"
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{"example 1", []int{1, 2, 5}, 11, 3},
		{"example 2", []int{2}, 3, -1},
		{"zero amount", []int{1}, 0, 0},
		{"single coin exact", []int{1}, 1, 1},
		{"single coin repeated", []int{1}, 2, 2},
		{"greedy fails", []int{1, 3, 4}, 6, 2},
		{"unsorted coins", []int{5, 1, 2}, 11, 3},
		{"mixed", []int{2, 5, 10, 1}, 27, 4},
		{"coin bigger than amount", []int{5, 10}, 3, -1},
		{"unreachable with gcd", []int{4, 6}, 7, -1},
		{"big coins", []int{186, 419, 83, 408}, 6249, 20},
		{"huge coin value", []int{2147483647}, 2, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := slices.Clone(tt.coins)
			got := coinChange(input, tt.amount)

			if !slices.Equal(input, tt.coins) {
				t.Fatalf("input mutated: %v, was %v", input, tt.coins)
			}

			if got != tt.want {
				t.Fatalf("got %d, want %d", got, tt.want)
			}
		})
	}
}
