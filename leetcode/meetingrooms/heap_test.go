package meetingrooms

import (
	"math/rand"
	"slices"
	"testing"
)

func drain(h *IntHeap) []int {
	var out []int
	for h.Len() > 0 {
		v, _ := h.Pop()
		out = append(out, v)
	}
	return out
}

func TestEmpty(t *testing.T) {
	h := New()
	if h.Len() != 0 {
		t.Fatalf("Len = %d, want 0", h.Len())
	}
	if v, ok := h.Peek(); ok {
		t.Fatalf("Peek on empty = (%d, true)", v)
	}
	if v, ok := h.Pop(); ok {
		t.Fatalf("Pop on empty = (%d, true)", v)
	}
}

func TestPushPeek(t *testing.T) {
	h := New()
	for _, x := range []int{5, 3, 8, 1, 9} {
		h.Push(x)
	}
	if h.Len() != 5 {
		t.Fatalf("Len = %d, want 5", h.Len())
	}
	if v, ok := h.Peek(); !ok || v != 1 {
		t.Fatalf("Peek = (%d, %v), want (1, true)", v, ok)
	}
	if h.Len() != 5 {
		t.Fatal("Peek must not remove elements")
	}
}

func TestPopOrder(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"mixed", []int{5, 3, 8, 1, 9, 2}, []int{1, 2, 3, 5, 8, 9}},
		{"duplicates", []int{2, 2, 1, 1, 3, 3}, []int{1, 1, 2, 2, 3, 3}},
		{"single", []int{42}, []int{42}},
		{"sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reversed", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"negative", []int{-1, -5, 0, 3}, []int{-5, -1, 0, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New()
			for _, x := range tt.in {
				h.Push(x)
			}
			if got := drain(h); !slices.Equal(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterleaved(t *testing.T) {
	h := New()
	h.Push(5)
	h.Push(2)
	if v, _ := h.Pop(); v != 2 {
		t.Fatalf("got %d, want 2", v)
	}
	h.Push(1)
	h.Push(7)
	for _, want := range []int{1, 5, 7} {
		if v, _ := h.Pop(); v != want {
			t.Fatalf("got %d, want %d", v, want)
		}
	}
	if _, ok := h.Pop(); ok {
		t.Fatal("heap should be empty")
	}
}

func TestFrom(t *testing.T) {
	in := []int{9, 4, 7, 1, 8, 2, 6, 3, 5}
	orig := slices.Clone(in)

	h := From(in)
	if h.Len() != len(in) {
		t.Fatalf("Len = %d, want %d", h.Len(), len(in))
	}
	if !slices.Equal(in, orig) {
		t.Fatal("From must not modify the input slice")
	}
	want := slices.Clone(in)
	slices.Sort(want)
	if got := drain(h); !slices.Equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestFromEmpty(t *testing.T) {
	h := From(nil)
	if h.Len() != 0 {
		t.Fatalf("Len = %d, want 0", h.Len())
	}
	if _, ok := h.Pop(); ok {
		t.Fatal("Pop on empty must return false")
	}
}

func TestRandomAgainstSort(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for round := 0; round < 200; round++ {
		in := make([]int, rng.Intn(200))
		for i := range in {
			in[i] = rng.Intn(50) - 25
		}
		want := slices.Clone(in)
		slices.Sort(want)

		h := New()
		for _, x := range in {
			h.Push(x)
		}
		if got := drain(h); !slices.Equal(got, want) {
			t.Fatalf("push round %d: got %v, want %v", round, got, want)
		}
		if got := drain(From(in)); !slices.Equal(got, want) {
			t.Fatalf("from round %d: got %v, want %v", round, got, want)
		}
	}
}
