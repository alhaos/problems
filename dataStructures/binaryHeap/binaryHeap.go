package binaryHeap

import "errors"

var ErrEmptyHeap = errors.New("heap is empty")

// Интерфейс двоичной кучи
type BinaryHeap interface {
	Push(n int)
	Pop() (int, error)
	Peek() (int, error)
	Len() int
	IsEmpty() bool
}

// binaryHeap структура двоичной кучи
type binaryHeap struct {
	data []int
}

// NewBinaryHeap возвращает новый экземпляр структуры BinaryHeap
func NewBinaryHeap() BinaryHeap {
	return &binaryHeap{
		data: make([]int, 0, 16),
	}
}

// Push добавляет элемент в двоичную кучу
func (h *binaryHeap) Push(n int) {

	// Добавляем в конец массива элемент
	h.data = append(h.data, n)

	// Фиксируем индекс элемента для просеивания вверх
	currentIndex := len(h.data) - 1

	// Цикл к корню кучи
	for currentIndex > 0 {

		// Находим индекс родительского элемента
		parentIndex := (currentIndex - 1) / 2

		// Если родительский элемент больше
		if h.data[parentIndex] > h.data[currentIndex] {

			// Меняем местами с текущим
			h.data[parentIndex], h.data[currentIndex] = h.data[currentIndex], h.data[parentIndex]

			// Изменяем индекс текущего элемента на индекс родительского элемента
			currentIndex = parentIndex
		} else { // Если родительский элемент не больше
			break // прерываем цикл
		}
	}
}

// Pop извлекает корневой элемент из бинарной кучи
func (h *binaryHeap) Pop() (int, error) {

	// Если куча пуста возвращаем 0 и ошибку
	if len(h.data) == 0 {
		return 0, ErrEmptyHeap
	}

	// Забираем элемент для возврата
	returnValue := h.data[0]

	// На его место вставляем крайний элемент
	h.data[0] = h.data[len(h.data)-1]

	// Обрезаем слайс
	h.data = h.data[:len(h.data)-1]

	// Устанавливаем 0 как индекс просматриваемого элемента
	currentIndex := 0

	// Просеиваем вниз
	for {
		leftChildIndex := 2*currentIndex + 1
		if leftChildIndex >= len(h.data) {
			break
		}
		rightChildIndex := 2*currentIndex + 2

		// выбираем наименьшего ребёнка
		smallestIndex := leftChildIndex
		if rightChildIndex < len(h.data) && h.data[rightChildIndex] < h.data[leftChildIndex] {
			smallestIndex = rightChildIndex
		}

		// если текущий элемент уже меньше — стоп
		if h.data[currentIndex] <= h.data[smallestIndex] {
			break
		}

		h.data[currentIndex], h.data[smallestIndex] = h.data[smallestIndex], h.data[currentIndex]
		currentIndex = smallestIndex
	}

	return returnValue, nil
}

// Peek Возвращает элемент в корне кучи
func (h *binaryHeap) Peek() (int, error) {
	if len(h.data) == 0 {
		return 0, ErrEmptyHeap
	}
	return h.data[0], nil
}

// Len возвращает длину двоичной кучи
func (h *binaryHeap) Len() int {
	return len(h.data)
}

// IsEmpty возвращает true если куча пуста, иначе false
func (h *binaryHeap) IsEmpty() bool {
	return len(h.data) == 0
}
