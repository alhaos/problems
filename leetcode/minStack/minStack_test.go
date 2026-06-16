package minStack

import "testing"

func TestMinStack(t *testing.T) {
	s := NewMinStack()

	s.Push(-2)
	s.Push(0)
	s.Push(-3)

	if got := s.GetMin(); got != -3 {
		t.Errorf("GetMin() = %d, want -3", got)
	}

	s.Pop()

	if got := s.Top(); got != 0 {
		t.Errorf("Top() = %d, want 0", got)
	}

	if got := s.GetMin(); got != -2 {
		t.Errorf("GetMin() = %d, want -2", got)
	}
}

func TestMinStackSingleElement(t *testing.T) {
	s := NewMinStack()
	s.Push(5)

	if got := s.Top(); got != 5 {
		t.Errorf("Top() = %d, want 5", got)
	}
	if got := s.GetMin(); got != 5 {
		t.Errorf("GetMin() = %d, want 5", got)
	}

	s.Pop()
	s.Push(3)

	if got := s.GetMin(); got != 3 {
		t.Errorf("GetMin() = %d, want 3", got)
	}
}

func TestMinStackMinAfterPop(t *testing.T) {
	s := NewMinStack()
	s.Push(2)
	s.Push(1)
	s.Push(3)

	if got := s.GetMin(); got != 1 {
		t.Errorf("GetMin() = %d, want 1", got)
	}

	s.Pop() // убираем 3, мин должен остаться 1
	if got := s.GetMin(); got != 1 {
		t.Errorf("GetMin() = %d, want 1", got)
	}

	s.Pop() // убираем 1, мин должен стать 2
	if got := s.GetMin(); got != 2 {
		t.Errorf("GetMin() = %d, want 2", got)
	}
}
