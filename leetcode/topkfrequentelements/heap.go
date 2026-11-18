package topkfrequentelements

import "container/heap"

type item struct {
	num int
	cnt int
}

// minHeap — min-heap по cnt.
type minHeap []item

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].cnt < h[j].cnt
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x any) {

	it := x.(item)
	*h = append(*h, it)

}

func (h *minHeap) Pop() any { return nil } // TODO

var _ heap.Interface = (*minHeap)(nil)

func topKFrequentHeap(nums []int, k int) []int {
	// TODO
	return nil
}
