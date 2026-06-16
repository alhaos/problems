package minStack

// MinStack структура стека с поддержкой минимального элемента
type MinStack struct {
	// items слайс элементов стека
	items []int
	// min слайс элементов минимумов нужен для обеспечения получения минимального элемента
	// за O(1) позволяет узнать какой минимум был до того как мы удалили минимальный элемент
	min []int
}

// NewMinStack конструктор для MinStack
func NewMinStack() *MinStack {
	return &MinStack{}
}

// Push добавляет элемент в стек
// обновляет минимум
func (s *MinStack) Push(val int) {
	s.items = append(s.items, val)
	if len(s.min) == 0 {
		s.min = append(s.min, val)
	} else {
		s.min = append(s.min, min(s.min[len(s.min)-1], val))
	}
}

func (s *MinStack) Pop() {
	// Удаляем крайний элемент
	s.items = s.items[:len(s.items)-1]
	// Обновляем минимум
	s.min = s.min[:len(s.min)-1]
}

// Top возвращает верхний элемент стека
func (s *MinStack) Top() int {
	// Возвращаем крайний элемент слайса items
	return s.items[len(s.items)-1]
}

// GetMin возвращаем минимальный элемент стека
func (s *MinStack) GetMin() int {
	// Возвращаем крайний элемент слайса min
	return s.min[len(s.min)-1]
}
