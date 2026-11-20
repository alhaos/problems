package kthlargest

import "testing"

func TestKthLargest(t *testing.T) {
	tests := []struct {
		name string
		k    int
		nums []int
		adds []int
		want []int
	}{
		{"example", 3, []int{4, 5, 8, 2}, []int{3, 5, 10, 9, 4}, []int{4, 5, 5, 8, 8}},
		{"k=1 empty start", 1, nil, []int{-3, -2, -4, 0, 4}, []int{-3, -2, -2, 0, 4}},
		{"k=2 one initial", 2, []int{0}, []int{-1, 1, -2, -4, 3}, []int{-1, 0, 0, 0, 1}},
		{"duplicates", 2, []int{5, 5}, []int{5, 6, 5}, []int{5, 5, 5}},
		{"fewer than k initially", 3, []int{1, 2}, []int{3, 4, 5}, []int{1, 2, 3}},
		{"nothing changes", 2, []int{10, 20}, []int{1, 2, 3}, []int{10, 10, 10}},
		{"growing", 2, []int{1, 2}, []int{3, 4, 5}, []int{2, 3, 4}},
		{"negatives", 3, []int{-5, -1, -3}, []int{-2, -10, 0}, []int{-3, -3, -2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kl := NewKthLargest(tt.k, tt.nums)
			for i, v := range tt.adds {
				if got := kl.Add(v); got != tt.want[i] {
					t.Errorf("Add(%d) #%d: got %d, want %d", v, i, got, tt.want[i])
				}
			}
		})
	}
}
