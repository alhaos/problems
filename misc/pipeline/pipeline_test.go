package pipeline

import (
	"context"
	"sort"
	"testing"
	"time"
)

func TestPipeline_Basic(t *testing.T) {
	ctx := context.Background()
	out := Pipeline(ctx, "HE", "hello", "world", "help", "go")

	var got []string
	for s := range out {
		got = append(got, s)
	}

	want := []string{"HELLO", "HELP"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d: got %q, want %q", i, got[i], want[i])
		}
	}
}

func TestPipeline_NoMatch(t *testing.T) {
	ctx := context.Background()
	out := Pipeline(ctx, "ZZZ", "hello", "world", "help")

	var got []string
	for s := range out {
		got = append(got, s)
	}

	if len(got) != 0 {
		t.Errorf("expected empty result, got %v", got)
	}
}

func TestPipeline_EmptyInput(t *testing.T) {
	ctx := context.Background()
	out := Pipeline(ctx, "HE")

	var got []string
	for s := range out {
		got = append(got, s)
	}

	if len(got) != 0 {
		t.Errorf("expected empty result, got %v", got)
	}
}

func TestPipeline_CancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// большой input чтобы pipeline не успел завершиться сам
	words := make([]string, 1000)
	for i := range words {
		words[i] = "hello"
	}

	out := Pipeline(ctx, "HE", words...)

	// читаем несколько значений и отменяем
	<-out
	<-out
	cancel()

	// pipeline должен завершиться, канал закрыться
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
		t.Fatal("pipeline did not stop after context cancel")
	}
}

func TestPipeline_ChannelClosed(t *testing.T) {
	ctx := context.Background()
	out := Pipeline(ctx, "GO", "golang", "gopher", "hello")

	var got []string
	for s := range out {
		got = append(got, s)
	}

	sort.Strings(got)
	want := []string{"GOLANG", "GOPHER"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d: got %q, want %q", i, got[i], want[i])
		}
	}
}
