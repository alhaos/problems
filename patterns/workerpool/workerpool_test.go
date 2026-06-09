package workerpool_test

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/alhaos/problems/patterns/workerpool"
)

func TestWorkerPool_Basic(t *testing.T) {
	ctx := context.Background()

	jobs := make(chan int, 5)
	for _, n := range []int{1, 2, 3, 4, 5} {
		jobs <- n
	}
	close(jobs)

	results := workerpool.WorkerPool(ctx, 3, jobs)

	var got []int
	for r := range results {
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

func TestWorkerPool_EmptyJobs(t *testing.T) {
	ctx := context.Background()

	jobs := make(chan int)
	close(jobs)

	results := workerpool.WorkerPool(ctx, 3, jobs)

	var got []int
	for r := range results {
		got = append(got, r)
	}

	if len(got) != 0 {
		t.Errorf("expected empty result, got %v", got)
	}
}

func TestWorkerPool_SingleWorker(t *testing.T) {
	ctx := context.Background()

	jobs := make(chan int, 3)
	for _, n := range []int{2, 3, 4} {
		jobs <- n
	}
	close(jobs)

	results := workerpool.WorkerPool(ctx, 1, jobs)

	var got []int
	for r := range results {
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

func TestWorkerPool_CancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	jobs := make(chan int, 1000)
	for i := range 1000 {
		jobs <- i
	}
	close(jobs)

	results := workerpool.WorkerPool(ctx, 5, jobs)

	<-results
	<-results
	cancel()

	done := make(chan struct{})
	go func() {
		for range results {
		}
		close(done)
	}()

	select {
	case <-done:
		// ok
	case <-time.After(time.Second):
		t.Fatal("worker pool did not stop after context cancel")
	}
}

func TestWorkerPool_MoreWorkersThanJobs(t *testing.T) {
	ctx := context.Background()

	jobs := make(chan int, 2)
	jobs <- 3
	jobs <- 4
	close(jobs)

	results := workerpool.WorkerPool(ctx, 10, jobs)

	var got []int
	for r := range results {
		got = append(got, r)
	}

	sort.Ints(got)
	want := []int{9, 16}

	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d: got %d, want %d", i, got[i], want[i])
		}
	}
}
