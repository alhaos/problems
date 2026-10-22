package robber

import (
	"math/rand"
	"testing"
)

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"single", []int{5}, 5},
		{"two", []int{1, 2}, 2},
		{"two reversed", []int{2, 1}, 2},
		{"example1", []int{1, 2, 3, 1}, 4},
		{"example2", []int{2, 7, 9, 3, 1}, 12},
		{"all zeros", []int{0, 0, 0}, 0},
		{"all same", []int{4, 4, 4, 4}, 8},
		{"skip two in a row", []int{5, 1, 1, 5}, 10},
		{"big in middle", []int{1, 100, 1}, 100},
		{"greedy fails", []int{2, 3, 2}, 4},
		{"greedy fails 2", []int{6, 7, 1, 30, 8, 2, 4}, 41},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.nums); got != tt.want {
				t.Errorf("rob(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

// brute перебирает все допустимые наборы домов.
func brute(nums []int, i int) int {
	if i >= len(nums) {
		return 0
	}
	take := nums[i] + brute(nums, i+2)
	skip := brute(nums, i+1)
	return max(take, skip)
}

func TestRobRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < 20000; i++ {
		nums := make([]int, rng.Intn(14)+1)
		for j := range nums {
			nums[j] = rng.Intn(10)
		}

		if got, want := rob(nums), brute(nums, 0); got != want {
			t.Fatalf("nums=%v got=%d want=%d", nums, got, want)
		}
	}
}
