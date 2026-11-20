package topk

import "container/heap"

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	fm := &freqHeap{freq: m}

	for key, _ := range m {
		heap.Push(fm, key)
		if fm.Len() > k {
			heap.Pop(fm)
		}
	}

	return fm.items
}
