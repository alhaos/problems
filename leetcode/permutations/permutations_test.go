package permutations

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int // ожидаемое количество перестановок
	}{
		{"one element", []int{1}, 1},
		{"two elements", []int{0, 1}, 2},
		{"three elements", []int{1, 2, 3}, 6},
		{"negatives", []int{-1, 0, 1}, 6},
		{"four elements", []int{4, 3, 2, 1}, 24},
		{"six elements", []int{1, 2, 3, 4, 5, 6}, 720},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)

			if len(got) != tt.want {
				t.Fatalf("got %d permutations, want %d", len(got), tt.want)
			}

			seen := make(map[string]bool, len(got))
			for _, p := range got {
				if len(p) != len(tt.nums) {
					t.Fatalf("permutation %v has wrong length", p)
				}
				key := fmt.Sprint(p)
				if seen[key] {
					t.Fatalf("duplicate permutation %v", p)
				}
				seen[key] = true
			}

			// каждая перестановка состоит ровно из элементов nums
			want := make(map[int]int, len(tt.nums))
			for _, n := range tt.nums {
				want[n]++
			}
			for _, p := range got {
				cnt := make(map[int]int, len(p))
				for _, n := range p {
					cnt[n]++
				}
				for k, v := range want {
					if cnt[k] != v {
						t.Fatalf("permutation %v has wrong elements", p)
					}
				}
			}
		})
	}
}
