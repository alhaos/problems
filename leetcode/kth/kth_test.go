package kth

import (
	"math/rand"
	"slices"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"single", []int{1}, 1, 1},
		{"example1", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"example2", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{"k equals 1", []int{7, 3, 9, 1}, 1, 9},
		{"k equals n", []int{7, 3, 9, 1}, 4, 1},
		{"all same", []int{5, 5, 5, 5}, 2, 5},
		{"negatives", []int{-1, -5, -3, -2}, 2, -2},
		{"duplicates count", []int{1, 2, 2, 3}, 3, 2},
		{"sorted asc", []int{1, 2, 3, 4, 5}, 2, 4},
		{"sorted desc", []int{5, 4, 3, 2, 1}, 4, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(slices.Clone(tt.nums), tt.k); got != tt.want {
				t.Errorf("findKthLargest(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func referenceKthLargest(nums []int, k int) int {
	s := slices.Clone(nums)
	slices.Sort(s)
	return s[len(s)-k]
}

func TestFindKthLargestRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < 100000; i++ {
		nums := make([]int, rng.Intn(15)+1)
		for j := range nums {
			nums[j] = rng.Intn(6) - 3
		}
		k := rng.Intn(len(nums)) + 1

		got := findKthLargest(slices.Clone(nums), k)
		want := referenceKthLargest(nums, k)
		if got != want {
			t.Fatalf("nums=%v k=%d got=%d want=%d", nums, k, got, want)
		}
	}
}
