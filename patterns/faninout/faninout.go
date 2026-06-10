package faninout

import (
	"context"
	"sync"
)

func FanOutIn(ctx context.Context, numWorkers int, in <-chan int) <-chan int {
	channels := FanOut(ctx, numWorkers, in)

	// каждый канал обрабатываем — возводим в квадрат
	results := make([]<-chan int, numWorkers)
	for i, ch := range channels {
		results[i] = worker(ctx, ch)
	}

	return FanIn(ctx, results...)
}

func worker(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case out <- v * v:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func FanIn(ctx context.Context, channels ...<-chan int) <-chan int {

	out := make(chan int)

	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				select {
				case out <- v:
				case <-ctx.Done():
					return
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func FanOut(ctx context.Context, numWorkers int, in <-chan int) []<-chan int {
	channels := make([]chan int, numWorkers)
	for i := range numWorkers {
		channels[i] = make(chan int)
	}

	go func() {
		defer func() {
			for _, ch := range channels {
				close(ch)
			}
		}()
		i := 0
		for v := range in {
			select {
			case channels[i] <- v:
				i = (i + 1) % numWorkers
			case <-ctx.Done():
				return
			}
		}
	}()

	result := make([]<-chan int, numWorkers)
	for i, ch := range channels {
		result[i] = ch
	}
	return result
}
