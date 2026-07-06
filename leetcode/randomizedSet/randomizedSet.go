package randomizedSet

import "math/rand"

// RandomizedSet реализует набор целых чисел с операциями Insert,
// Remove и GetRandom выполняемыми за O(1) по скорости
type RandomizedSet struct {
	// Карта хранит элементы в виде ключей и индексы в массиве
	// values в качестве значений
	data map[int]int
	// Для обеспечения выполнения операции GetRandom за O(1)
	// приходится жертвовать памятью и храните еще один набор
	// элементов в виде массива
	values []int
}

// NewRandomizedSet конструктор для структуры RandomizedSet
func NewRandomizedSet() RandomizedSet {
	return RandomizedSet{
		data:   make(map[int]int),
		values: []int{},
	}
}

// Insert добавляет элемент в RandomizedSet возвращает true в случае успеха
// false в случае если элемент уже присутствует в наборе
func (s *RandomizedSet) Insert(val int) bool {

	// Проверяем наличие элемента в карте элементов делается за O(1)
	if _, exist := s.data[val]; exist {
		// Если есть возвращаем false
		return false
	}

	// Если нет
	// Добавляем в массив сложность O(1)
	s.values = append(s.values, val)
	// Добавляем в карту сложность O(1)
	s.data[val] = len(s.values) - 1
	// Возвращаем true
	return true
}

// Remove удаляет элемент из набора
func (s *RandomizedSet) Remove(val int) bool {

	// Получаем индекс удаляемого элемента и флаг его наличия
	idx, exist := s.data[val]
	// Проверяем наличие элемента
	if !exist {
		// Если нет возвращаем false
		return false
	}

	// Если есть
	// Вычисляем lastIdx
	lastIdx := len(s.values) - 1

	// Получаем крайний в массиве элемент O(1)
	lastVal := s.values[lastIdx]

	// Меняем удаляемый элемент с крайним
	s.values[idx], s.values[lastIdx] = s.values[lastIdx], s.values[idx]

	// Усекаем массив на один элемент O(1)
	s.values = s.values[:len(s.values)-1]

	// Обновляем индекс для бывшего крайнего элемента o(1)
	s.data[lastVal] = idx

	// Удаляем элемент из карты O(1)
	delete(s.data, val)

	// Возвращаем true
	return true
}

// Возвращает случайный элемент
func (s *RandomizedSet) GetRandom() int {
	// Возвращаем значение получив случайный индекс массива
	return s.values[rand.Intn(len(s.values))]
}
