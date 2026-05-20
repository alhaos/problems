package binaryHeap

import (
	"errors"
	"testing"
)

func TestPeekOnEmptyHeap(t *testing.T) {
	h := NewBinaryHeap()
	_, err := h.Peek()
	if err == nil {
		t.Error("expected error on empty heap")
	}
}

func TestPopOnEmptyHeap(t *testing.T) {
	h := NewBinaryHeap()
	_, err := h.Pop()
	if err == nil {
		t.Error("expected error on empty heap")
	}
}

func TestIsEmptyOnNewHeap(t *testing.T) {
	h := NewBinaryHeap()
	if !h.IsEmpty() {
		t.Error("expected new heap to be empty")
	}
}

func TestLenOnNewHeap(t *testing.T) {
	h := NewBinaryHeap()
	if h.Len() != 0 {
		t.Errorf("expected len 0, got %d", h.Len())
	}
}

func TestPushSingleElement(t *testing.T) {
	h := NewBinaryHeap()
	h.Push(42)

	if h.IsEmpty() {
		t.Error("expected heap to be non-empty after Push")
	}
	if h.Len() != 1 {
		t.Errorf("expected len 1, got %d", h.Len())
	}

	val, err := h.Peek()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 42 {
		t.Errorf("expected 42, got %d", val)
	}
}

func TestPushMultipleElementsPeekReturnsMin(t *testing.T) {
	h := NewBinaryHeap()
	h.Push(30)
	h.Push(10)
	h.Push(20)

	val, err := h.Peek()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 10 {
		t.Errorf("expected min 10, got %d", val)
	}
}

func TestPopReturnsSortedOrder(t *testing.T) {
	h := NewBinaryHeap()
	input := []int{50, 10, 40, 20, 30}
	for _, v := range input {
		h.Push(v)
	}

	expected := []int{10, 20, 30, 40, 50}
	for _, want := range expected {
		got, err := h.Pop()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	}
}

func TestPopDecreasesLen(t *testing.T) {
	h := NewBinaryHeap()
	h.Push(1)
	h.Push(2)
	h.Push(3)

	h.Pop()
	if h.Len() != 2 {
		t.Errorf("expected len 2, got %d", h.Len())
	}
}

func TestIsEmptyAfterPopAll(t *testing.T) {
	h := NewBinaryHeap()
	h.Push(1)
	h.Pop()

	if !h.IsEmpty() {
		t.Error("expected heap to be empty after popping all elements")
	}
}

func TestPeekDoesNotRemoveElement(t *testing.T) {
	h := NewBinaryHeap()
	h.Push(5)

	h.Peek()
	if h.Len() != 1 {
		t.Errorf("expected len 1 after Peek, got %d", h.Len())
	}
}

func TestErrorType(t *testing.T) {
	h := NewBinaryHeap()
	_, err := h.Pop()
	if !errors.Is(err, ErrEmptyHeap) {
		t.Errorf("expected ErrEmptyHeap, got %v", err)
	}
}
