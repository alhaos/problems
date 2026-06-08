package pipeline_context

import (
	"context"
	"testing"
	"time"
)

// Тест 2: Processor удваивает значения
func TestProcessorDoublesValues(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := make(chan int, 3)
	in <- 1
	in <- 5
	in <- 10
	close(in)

	proc := Processor(ctx, in)

	expected := []int{2, 10, 20}
	for _, want := range expected {
		got, ok := <-proc
		if !ok {
			t.Fatal("channel closed too early")
		}
		if got != want {
			t.Fatalf("expected %d, got %d", want, got)
		}
	}
}

// Тест 3: пайплайн останавливается по отмене контекста
func TestPipelineStopsOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	gen := Generator(ctx)
	proc := Processor(ctx, gen)

	// читаем немного
	for i := 0; i < 5; i++ {
		<-proc
	}

	cancel()

	// даём горутинам время завершиться
	time.Sleep(100 * time.Millisecond)

	// канал должен быть закрыт
	select {
	case _, ok := <-proc:
		if ok {
			t.Fatal("expected channel to be closed after cancel")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("channel did not close after cancel — possible goroutine leak")
	}
}

// Тест 4: нет горутин-призраков после отмены
func TestNoGoroutineLeakOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	gen := Generator(ctx)
	proc := Processor(ctx, gen)

	<-proc
	cancel()

	// ждём завершения горутин
	time.Sleep(200 * time.Millisecond)

	// proc должен закрыться — дренируем остаток
	deadline := time.After(500 * time.Millisecond)
	for {
		select {
		case _, ok := <-proc:
			if !ok {
				return // всё хорошо
			}
		case <-deadline:
			t.Fatal("goroutine leak: channel still open after cancel")
		}
	}
}
