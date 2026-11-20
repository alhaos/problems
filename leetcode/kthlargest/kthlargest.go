package kthlargest

import "container/heap"

type minHeap []int

// value receiver
func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// pointer receiver
func (h *minHeap) Push(x any) { *h = append((*h), x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	l := old.Len()
	x := old[l-1]
	*h = old[:l-1]
	return x
}

type KthLargest struct {
	k int
	h minHeap
}

func NewKthLargest(k int, nums []int) *KthLargest {
	var h minHeap
	for _, n := range nums {
		heap.Push(&h, n)
		if h.Len() > k {
			heap.Pop(&h)
		}
	}
	return &KthLargest{k: k, h: h}
}

func (kl *KthLargest) Add(val int) int {

	heap.Push(&kl.h, val)

	if kl.h.Len() > kl.k {
		heap.Pop(&kl.h)
	}

	return kl.h[0]
}
