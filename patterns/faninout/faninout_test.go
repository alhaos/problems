package faninout

import (
	"context"
	"sort"
	"testing"
	"time"
)

func TestFanOutIn_Basic(t *testing.T) {
	ctx := context.Background()

	in := make(chan int, 5)
	for _, n := range []int{1, 2, 3, 4, 5} {
		in <- n
	}
	close(in)

	out := FanOutIn(ctx, 3, in)

	var got []int
	for r := range out {
		got = append(got, r)
	}

	sort.Ints(got)
	want := []int{1, 4, 9, 16, 25}

	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d: got %d, want %d", i, got[i], want[i])
		}
	}
}

func TestFanOutIn_EmptyInput(t *testing.T) {
	ctx := context.Background()

	in := make(chan int)
	close(in)

	out := FanOutIn(ctx, 3, in)

	var got []int
	for r := range out {
		got = append(got, r)
	}

	if len(got) != 0 {
		t.Errorf("expected empty result, got %v", got)
	}
}

func TestFanOutIn_SingleWorker(t *testing.T) {
	ctx := context.Background()

	in := make(chan int, 3)
	for _, n := range []int{2, 3, 4} {
		in <- n
	}
	close(in)

	out := FanOutIn(ctx, 1, in)

	var got []int
	for r := range out {
		got = append(got, r)
	}

	sort.Ints(got)
	want := []int{4, 9, 16}

	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d: got %d, want %d", i, got[i], want[i])
		}
	}
}

func TestFanOutIn_CancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	in := make(chan int, 1000)
	for i := range 1000 {
		in <- i
	}
	close(in)

	out := FanOutIn(ctx, 5, in)

	<-out
	<-out
	cancel()

	done := make(chan struct{})
	go func() {
		for range out {
		}
		close(done)
	}()

	select {
	case <-done:
		// ok
	case <-time.After(time.Second):
		t.Fatal("faninout did not stop after context cancel")
	}
}

func TestFanIn_MergesAllChannels(t *testing.T) {
	ctx := context.Background()

	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	ch3 := make(chan int, 2)

	ch1 <- 1
	ch1 <- 2
	ch2 <- 3
	ch2 <- 4
	ch3 <- 5
	ch3 <- 6

	close(ch1)
	close(ch2)
	close(ch3)

	out := FanIn(ctx, ch1, ch2, ch3)

	var got []int
	for r := range out {
		got = append(got, r)
	}

	sort.Ints(got)
	want := []int{1, 2, 3, 4, 5, 6}

	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d: got %d, want %d", i, got[i], want[i])
		}
	}
}

func TestFanOut_RoundRobin(t *testing.T) {
	ctx := context.Background()

	in := make(chan int, 6)
	for _, n := range []int{1, 2, 3, 4, 5, 6} {
		in <- n
	}
	close(in)

	channels := FanOut(ctx, 3, in)

	if len(channels) != 3 {
		t.Fatalf("expected 3 channels, got %d", len(channels))
	}

	var got []int
	for _, ch := range channels {
		for v := range ch {
			got = append(got, v)
		}
	}

	sort.Ints(got)
	want := []int{1, 2, 3, 4, 5, 6}

	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d: got %d, want %d", i, got[i], want[i])
		}
	}
}
