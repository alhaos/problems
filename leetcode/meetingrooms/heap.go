package meetingrooms

// IntHeap is a binary min-heap of ints: the smallest element is on top.
type IntHeap struct {
	data []int
}

// New creates an empty heap.
func New() *IntHeap {
	return &IntHeap{}
}

// Len returns the number of elements.
func (h *IntHeap) Len() int {
	return len(h.data)
}

// Push adds x to the heap. O(log n).
func (h *IntHeap) Push(x int) {
	h.data = append(h.data, x)
	siftUp(h.data, len(h.data)-1)
}

func siftUp(data []int, i int) {
	for {
		parent := (i - 1) / 2
		if data[i] < data[parent] {
			data[i], data[parent] = data[parent], data[i]
			i = parent
			continue
		}
		return
	}
}

// Peek returns the smallest element without removing it.
// Returns (0, false) if the heap is empty. O(1).
func (h *IntHeap) Peek() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	return h.data[0], true
}

// Pop removes and returns the smallest element.
// Returns (0, false) if the heap is empty. O(log n).
func (h *IntHeap) Pop() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}

	result := h.data[0]

	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	siftDown(h.data, 0)

	return result, true

}

func siftDown(data []int, i int) {
	for {
		left := i*2 + 1
		right := i*2 + 2

		if left > len(data)-1 {
			return
		}

		child := left
		if right < len(data) && data[right] < data[left] {
			child = right
		}

		if data[i] <= data[child] {
			return
		}

		data[i], data[child] = data[child], data[i]

		i = child
	}
}

func From(items []int) *IntHeap {
	heap := IntHeap{}

	for _, item := range items {
		heap.Push(item)
	}

	return &heap
}
