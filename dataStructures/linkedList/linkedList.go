package linkedlist

import (
	"fmt"
)

// compile-time проверка что LinkedList реализует List
// Это хитрый трюк Go — присваивание в _ (blank identifier)
// которое компилятор проверяет но не выполняет в runtime.
// компилятор видит присваивание и проверяет что правая часть
// реализует интерфейс левой части. Если какой-то метод не реализован — ошибка компиляции
var _ List[struct{}] = (*LinkedList[struct{}])(nil)

// List — интерфейс односвязного списка
type List[T comparable] interface {
	// PushFront добавляет элемент в начало списка — O(1)
	PushFront(val T)

	// PushBack добавляет элемент в конец списка — O(n)
	PushBack(val T)

	// PopFront удаляет и возвращает первый элемент — O(1)
	// Возвращает false если список пуст
	PopFront() (T, bool)

	// PopBack удаляет и возвращает последний элемент — O(n)
	// Возвращает false если список пуст
	PopBack() (T, bool)

	// PeekFront возвращает первый элемент без удаления — O(1)
	PeekFront() (T, bool)

	// PeekBack возвращает последний элемент без удаления — O(n)
	PeekBack() (T, bool)

	// Insert вставляет элемент по индексу — O(n)
	Insert(index int, val T) error

	// Remove удаляет элемент по индексу — O(n)
	Remove(index int) (T, error)

	// Len возвращает количество элементов — O(1)
	Len() int

	// Contains проверяет наличие элемента — O(n)
	Contains(val T) bool

	// ToSlice возвращает все элементы в виде среза — O(n)
	ToSlice() []T
}

// Node — узел односвязного списка
type Node[T comparable] struct {
	// Значение для хранения
	val T
	// Ссылка на последующий элемент
	next *Node[T]
}

// LinkedList — односвязный список
type LinkedList[T comparable] struct {
	// Ссылка на головной элемент
	head *Node[T]
	// Ссылка на хвостовой элемент
	tail *Node[T]
	// Длинна списка
	len int
}

// Вставляем значение в начало списка
func (l *LinkedList[T]) PushFront(val T) {

	// Инициируем новый элемент списка
	node := Node[T]{
		val: val,
	}

	// Ошибок мы не предполагаем поэтому счетчик увеличиваем сразу
	l.len++

	// Если список пустой
	if l.head == nil {

		// Устанавливаем в head ссылку на инициированный элемент
		l.head = &node

		// Устанавливаем в tail ссылку на инициированный элемент
		l.tail = &node

		// Закончили упражнение
		return
	}

	// Если код дошел до сюда список не пуст
	// Полю next нового элемента присваиваем ссылку на текущий элемент head
	node.next = l.head

	// Заменяем элемент head новым элементом
	l.head = &node
}

// Вставляем значение в конец списка
func (l *LinkedList[T]) PushBack(val T) {

	// Увеличиваем счетчик элементов
	l.len++

	// Инициируем новый элемент списка
	node := Node[T]{
		val: val,
	}

	// Если список пустой
	if l.head == nil {

		// Устанавливаем в head ссылку на инициированный элемент
		l.head = &node

		// Устанавливаем в tail ссылку на инициированный элемент
		l.tail = &node

		// Закончили упражнение
		return
	}

	// Значению поля next хвостового элемента присваиваем ссылку на новый элемент
	l.tail.next = &node

	// Значению поля tail списка присваиваем ссылку на новый хвостовой элемент
	l.tail = &node

}

// Извлекаем значение головного элемента
func (l *LinkedList[T]) PopFront() (T, bool) {

	// Если список пуст возвращаем Т, false
	if l.head == nil {
		var zero T
		return zero, false
	}

	// Если код дошел сюда то список не пуст
	value := l.head.val

	// Уменьшаем счетчик элементов
	l.len--

	// Если это был крайний элемент
	if l.head.next == nil {
		// В head устанавливаем ссылку на nil
		l.head = nil
		// В tail устанавливаем ссылку на nil
		l.tail = nil
	} else { // В этом случае есть еще элементы
		// Меняем ссылку головного элемента списка на ссылку на следующий элемент
		l.head = l.head.next
	}

	// Возвращаем значение и true
	return value, true
}

// Извлекаем значение хвостового элемента
func (l *LinkedList[T]) PopBack() (T, bool) {
	// Если список пуст возвращаем Т, false
	if l.head == nil {
		var zero T
		return zero, false
	}

	// Если код дошел сюда то список не пуст
	value := l.tail.val

	// Уменьшаем счетчик элементов
	l.len--

	// Если в списке один элемент
	if l.head.next == nil {
		// Удаляем головной элемент
		l.head = nil
		// Удаляем хвостовой элемент
		l.tail = nil
		// Возвращаем значение и true
		return value, true
	}

	// Тут надо перебрать весь список чтобы получить адрес предпоследнего элемента
	// ересь конечно, но оставлю этот метод

	// currentNode содержит текущий элемент при обходе
	// начинаем с головного элемента
	currentNode := l.head

	// обходим список пока у следующего за текущим элементом поле next не окажется пустым
	for currentNode.next.next != nil {
		// Присваиваем текущему элементу обхода следующий элемент
		currentNode = currentNode.next
	}

	// Когда цикл будет пройден до конца currentNode будет содержать предпоследний элемент,
	// Удаляем в нем ссылку на последний элемент
	currentNode.next = nil

	// Присваиваем его хвостовому элементу
	l.tail = currentNode

	// Возвращаем значение и true
	return value, true
}

// PeekFront возвращает значение головного элемента не извлекая его
func (l *LinkedList[T]) PeekFront() (T, bool) {

	// Если список пуст возвращаем Т, false
	if l.head == nil {
		var zero T
		return zero, false
	}

	// Возвращаем значение головного элемента
	return l.head.val, true
}

// PeekBack возвращает значение хвостового элемента не извлекая его
func (l *LinkedList[T]) PeekBack() (T, bool) {

	// Если список пуст возвращаем Т, false
	if l.head == nil {
		var zero T
		return zero, false
	}

	// Возвращаем значение хвостового элемента
	return l.tail.val, true
}

// Insert вставляет значение на определенный индекс списка
func (l *LinkedList[T]) Insert(index int, val T) error {

	// Проверяем что индекс вписывается в диапазон существующих значений
	if index < 0 || index > l.len {
		return ErrIndexOutOfRange
	}

	// Если индекс 0 пользуемся методом PushFront
	if index == 0 {
		l.PushFront(val)
		return nil
	}

	// если индекс равен длине списка то используем PushBack
	if index == l.len {
		l.PushBack(val)
		return nil
	}

	// Если код дошел до сюда то индекс внутри списка
	// Получаем предыдущий элемент
	currentNode := l.nodeAt(index - 1)

	// Создаем новый элемент
	newNode := Node[T]{
		val:  val,
		next: currentNode.next, // тут ссылка на элемент который нашли при обходе
	}

	// Меняем ссылку
	currentNode.next = &newNode

	// Увеличиваем счетчик
	l.len++

	// Возвращаем nil
	return nil
}

// Remove удаляет элемент из списка
func (l *LinkedList[T]) Remove(index int) (T, error) {

	// Проверяем корректность индекса
	if index < 0 || index >= l.len {
		var zero T
		return zero, ErrIndexOutOfRange
	}

	// При индексе 0 выполняем PopFront
	if index == 0 {
		val, _ := l.PopFront()
		return val, nil
	}

	// При хвостовом индексе удаляем PopBack
	if index == l.len-1 {
		val, _ := l.PopBack()
		return val, nil
	}

	// В списке больше одного элемента, индекс в середине
	prevNode := l.nodeAt(index - 1)
	val := prevNode.next.val
	prevNode.next = prevNode.next.next
	l.len--
	return val, nil
}

// nodeAt возвращает узел по индексу
func (l *LinkedList[T]) nodeAt(index int) *Node[T] {
	currentNode := l.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	return currentNode
}

// Len возвращает длину списка
func (l *LinkedList[T]) Len() int {
	return l.len
}

func (l *LinkedList[T]) Contains(val T) bool {
	if l.len == 0 {
		return false
	}
	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode.val == val {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) ToSlice() []T {

	slice := make([]T, 0, l.len)

	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		slice = append(slice, currentNode.val)
	}

	return slice
}

// sentinel error-ы
var (
	ErrEmptyList       = fmt.Errorf("linkedlist: list is empty")
	ErrIndexOutOfRange = fmt.Errorf("linkedlist: index out of range")
)
