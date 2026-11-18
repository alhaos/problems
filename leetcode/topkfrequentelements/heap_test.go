package topkfrequentelements

import (
	"container/heap"
	"slices"
	"testing"
)

func TestMinHeapOrder(t *testing.T) {
	h := &minHeap{}
	for _, it := range []item{{1, 5}, {2, 1}, {3, 3}, {4, 2}, {5, 4}} {
		heap.Push(h, it)
	}

	var got []int
	for h.Len() > 0 {
		got = append(got, heap.Pop(h).(item).cnt)
	}

	want := []int{1, 2, 3, 4, 5}
	if !slices.Equal(got, want) {
		t.Errorf("pop order by cnt = %v, want %v", got, want)
	}
}

func TestTopKFrequentHeap(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"example 1", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"example 2", []int{1}, 1, []int{1}},
		{"k = 1", []int{4, 4, 4, 5, 5, 6}, 1, []int{4}},
		{"k = all unique", []int{1, 2, 3, 1, 2, 1}, 3, []int{1, 2, 3}},
		{"negatives", []int{-1, -1, -2, -2, -2, 3}, 2, []int{-2, -1}},
		{"several in one bucket", []int{1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"all same", []int{7, 7, 7, 7}, 1, []int{7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequentHeap(slices.Clone(tt.nums), tt.k)
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("topKFrequentHeap(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
