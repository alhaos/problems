package laststone

import "container/heap"

type stoneHeap []int

func (h stoneHeap) Len() int           { return len(h) }
func (h stoneHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h stoneHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *stoneHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *stoneHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {

	h := stoneHeap(stones)
	heap.Init(&h)

	for h.Len() > 1 {
		first, second := heap.Pop(&h).(int), heap.Pop(&h).(int)
		if first != second {
			heap.Push(&h, first-second)
		}
	}

	if h.Len() == 0 {
		return 0
	}

	return h[0]
}
