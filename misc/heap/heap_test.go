package heap

import (
	"math/rand"
	"slices"
	"testing"
)

func minLess(a, b int) bool { return a < b }
func maxLess(a, b int) bool { return a > b }

func drain(h *IntHeap) []int {
	var res []int
	for h.Len() > 0 {
		v, ok := h.Pop()
		if !ok {
			break
		}
		res = append(res, v)
	}
	return res
}

func TestEmpty(t *testing.T) {
	h := New(minLess)
	if h.Len() != 0 {
		t.Errorf("Len = %d, want 0", h.Len())
	}
	if v, ok := h.Peek(); ok || v != 0 {
		t.Errorf("Peek = (%d, %v), want (0, false)", v, ok)
	}
	if v, ok := h.Pop(); ok || v != 0 {
		t.Errorf("Pop = (%d, %v), want (0, false)", v, ok)
	}
}

func TestPeekDoesNotRemove(t *testing.T) {
	h := New(minLess)
	h.Push(5)
	h.Push(2)
	h.Push(8)
	for i := 0; i < 3; i++ {
		if v, ok := h.Peek(); !ok || v != 2 {
			t.Fatalf("Peek = (%d, %v), want (2, true)", v, ok)
		}
	}
	if h.Len() != 3 {
		t.Errorf("Len = %d, want 3", h.Len())
	}
}

func TestPopOrder(t *testing.T) {
	inputs := map[string][]int{
		"single":     {4},
		"mixed":      {5, 3, 8, 1, 9, 2},
		"duplicates": {3, 1, 3, 2, 1, 2},
		"sorted":     {1, 2, 3, 4, 5},
		"reversed":   {5, 4, 3, 2, 1},
		"negative":   {-1, 3, 0, -5},
	}

	for name, in := range inputs {
		t.Run(name+"/min", func(t *testing.T) {
			h := New(minLess)
			for _, x := range in {
				h.Push(x)
			}
			want := slices.Clone(in)
			slices.Sort(want)
			if got := drain(h); !slices.Equal(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})

		t.Run(name+"/max", func(t *testing.T) {
			h := New(maxLess)
			for _, x := range in {
				h.Push(x)
			}
			want := slices.Clone(in)
			slices.Sort(want)
			slices.Reverse(want)
			if got := drain(h); !slices.Equal(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestInterleaved(t *testing.T) {
	h := New(minLess)
	h.Push(5)
	h.Push(3)
	if v, _ := h.Pop(); v != 3 {
		t.Fatalf("got %d, want 3", v)
	}
	h.Push(1)
	h.Push(4)
	if v, _ := h.Pop(); v != 1 {
		t.Fatalf("got %d, want 1", v)
	}
	if v, _ := h.Pop(); v != 4 {
		t.Fatalf("got %d, want 4", v)
	}
	if v, _ := h.Pop(); v != 5 {
		t.Fatalf("got %d, want 5", v)
	}
	if h.Len() != 0 {
		t.Errorf("Len = %d, want 0", h.Len())
	}
}

func TestFrom(t *testing.T) {
	in := []int{9, 4, 7, 1, 8, 2, 6, 3, 5}
	orig := slices.Clone(in)

	h := From(in, minLess)
	want := slices.Clone(in)
	slices.Sort(want)
	if got := drain(h); !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	// From не должна портить входной слайс
	if !slices.Equal(in, orig) {
		t.Errorf("input mutated: %v, was %v", in, orig)
	}
}

func TestFromEmpty(t *testing.T) {
	h := From(nil, minLess)
	if h.Len() != 0 {
		t.Errorf("Len = %d, want 0", h.Len())
	}
}

func TestRandomAgainstSort(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for round := 0; round < 2000; round++ {
		n := rng.Intn(60)
		in := make([]int, n)
		for i := range in {
			in[i] = rng.Intn(21) - 10
		}

		want := slices.Clone(in)
		slices.Sort(want)

		h := New(minLess)
		for _, x := range in {
			h.Push(x)
		}
		if got := drain(h); !slices.Equal(got, want) {
			t.Fatalf("min push round %d: got %v, want %v", round, got, want)
		}

		if got := drain(From(in, minLess)); !slices.Equal(got, want) {
			t.Fatalf("min from round %d: got %v, want %v", round, got, want)
		}

		slices.Reverse(want)
		if got := drain(From(in, maxLess)); !slices.Equal(got, want) {
			t.Fatalf("max from round %d: got %v, want %v", round, got, want)
		}
	}
}
