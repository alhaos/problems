package heap

type IntHeap struct {
	data []int
	less func(a, b int) bool
}

func New(less func(a, b int) bool) *IntHeap {
	return &IntHeap{less: less}
}

func From(items []int, less func(a, b int) bool) *IntHeap {

	h := New(less)

	for _, item := range items {
		h.Push(item)
	}

	return h
}

func (h *IntHeap) Len() int {
	return len(h.data)
}

func (h *IntHeap) Push(x int) {
	h.data = append(h.data, x)
	siftUp(h.data, len(h.data)-1, h.less)
}

func (h *IntHeap) Pop() (int, bool) {
	if h.Len() > 0 {
		result := h.data[0]
		h.data[0] = h.data[h.Len()-1]
		h.data = h.data[:h.Len()-1]
		siftDown(h.data, 0, h.less)
		return result, true
	}

	return 0, false
}

func (h *IntHeap) Peek() (int, bool) {
	if h.Len() > 0 {
		return h.data[0], true
	}
	return 0, false
}

func siftUp(data []int, i int, less func(a, b int) bool) {
	for i > 0 {
		parent := (i - 1) / 2

		if less(data[i], data[parent]) {
			data[parent], data[i] = data[i], data[parent]
			i = parent
			continue
		}

		return
	}

}

func siftDown(data []int, i int, less func(a, b int) bool) {

	for {

		left := i*2 + 1
		right := i*2 + 2

		// Если нет левого ребенка выходим просеивание окончено выходим из цикла
		if left > len(data)-1 {
			return
		}

		child := left

		// Если значение правого ребенка в пределах массива и правый элемент меньше левого
		// то ребенком для проверки выбирается правый
		if right < len(data) && less(data[right], data[left]) {
			child = right
		}

		// Если значение элемента с индексом ребенка не лучше чем значение элемента с индексом i
		// просеивание окончено выходим из цикла
		if !less(data[child], data[i]) {
			return
		}

		data[i], data[child] = data[child], data[i]

		i = child
	}
}
