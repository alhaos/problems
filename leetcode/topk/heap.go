package topk

type freqHeap struct {
	items []int
	freq  map[int]int
}

// эти три по сути для сортировки, value receiver
func (h freqHeap) Len() int {
	return len(h.items)
}

func (h freqHeap) Less(i, j int) bool {
	return h.freq[h.items[i]] < h.freq[h.items[j]]
}

func (h freqHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

// эти два меняют длину слайса, нужен pointer receiver
func (h *freqHeap) Push(x any) { h.items = append(h.items, x.(int)) }
func (h *freqHeap) Pop() any {
	n := len(h.items)
	x := h.items[n-1]
	h.items = h.items[:n-1]
	return x
}
