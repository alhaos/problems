package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushFrontAndToSlice(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushFront(3)
	l.PushFront(2)
	l.PushFront(1)
	assert.Equal(t, []int{1, 2, 3}, l.ToSlice())
}

func TestPushBack(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	assert.Equal(t, []int{1, 2, 3}, l.ToSlice())
}

func TestPopFront(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	val, ok := l.PopFront()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	assert.Equal(t, []int{2, 3}, l.ToSlice())
}

func TestPopFrontEmpty(t *testing.T) {
	l := &LinkedList[int]{}
	_, ok := l.PopFront()
	assert.False(t, ok)
}

func TestPopBack(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	val, ok := l.PopBack()
	assert.True(t, ok)
	assert.Equal(t, 3, val)
	assert.Equal(t, []int{1, 2}, l.ToSlice())
}

func TestPopBackEmpty(t *testing.T) {
	l := &LinkedList[int]{}
	_, ok := l.PopBack()
	assert.False(t, ok)
}

func TestPeekFront(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushBack(1)
	l.PushBack(2)

	val, ok := l.PeekFront()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	assert.Equal(t, 2, l.Len()) // peek не удаляет
}

func TestPeekBack(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushBack(1)
	l.PushBack(2)

	val, ok := l.PeekBack()
	assert.True(t, ok)
	assert.Equal(t, 2, val)
	assert.Equal(t, 2, l.Len()) // peek не удаляет
}

func TestInsert(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushBack(1)
	l.PushBack(3)

	err := l.Insert(1, 2) // [1, 2, 3]
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, l.ToSlice())
}

func TestInsertAtHead(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushBack(2)
	l.PushBack(3)

	err := l.Insert(0, 1) // [1, 2, 3]
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, l.ToSlice())
}

func TestInsertAtTail(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushBack(1)
	l.PushBack(2)

	err := l.Insert(2, 3) // [1, 2, 3]
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, l.ToSlice())
}

func TestInsertOutOfRange(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushBack(1)

	err := l.Insert(5, 99)
	assert.ErrorIs(t, err, ErrIndexOutOfRange)
}

func TestRemove(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	val, err := l.Remove(1) // удаляем 2
	assert.NoError(t, err)
	assert.Equal(t, 2, val)
	assert.Equal(t, []int{1, 3}, l.ToSlice())
}

func TestRemoveHead(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushBack(1)
	l.PushBack(2)

	val, err := l.Remove(0)
	assert.NoError(t, err)
	assert.Equal(t, 1, val)
	assert.Equal(t, []int{2}, l.ToSlice())
}

func TestRemoveOutOfRange(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushBack(1)

	_, err := l.Remove(5)
	assert.ErrorIs(t, err, ErrIndexOutOfRange)
}

func TestLen(t *testing.T) {
	l := &LinkedList[int]{}
	assert.Equal(t, 0, l.Len())
	l.PushBack(1)
	l.PushBack(2)
	assert.Equal(t, 2, l.Len())
	l.PopFront()
	assert.Equal(t, 1, l.Len())
}

func TestContains(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	assert.True(t, l.Contains(2))
	assert.False(t, l.Contains(99))
}

func TestToSliceEmpty(t *testing.T) {
	l := &LinkedList[int]{}
	assert.Equal(t, []int{}, l.ToSlice())
}
