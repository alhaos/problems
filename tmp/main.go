package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // min-куча
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	// 1. Добавление по одному
	h := &IntHeap{}
	heap.Push(h, 5)
	heap.Push(h, 2)
	heap.Push(h, 8)
	heap.Push(h, 1)

	fmt.Println("peek:", (*h)[0]) // вершина всегда на индексе 0
	fmt.Println("len:", h.Len())

	for h.Len() > 0 {
		fmt.Print(heap.Pop(h), " ")
	}
	fmt.Println()

	// 2. Построить кучу из готового слайса (heapify)
	h2 := &IntHeap{9, 4, 7, 1}
	heap.Init(h2)
	for h2.Len() > 0 {
		fmt.Print(heap.Pop(h2), " ")
	}
	fmt.Println()
}
