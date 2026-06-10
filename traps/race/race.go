package race

import (
	"sync"
	"sync/atomic"
)

func CounterWithRace(n int) int {
	var wg sync.WaitGroup
	var counter int
	for range n {
		wg.Go(
			func() {
				counter++
			},
		)
	}
	wg.Wait()
	return counter
}

// CounterWithAtomic функция для демонстрации решения состояния гонки
// при помощи примитива пакета atomic
func CounterWithAtomic(n int) int {
	var wg sync.WaitGroup
	var counter atomic.Int32
	for range n {
		wg.Go(
			func() {
				counter.Add(1)
			},
		)
	}
	wg.Wait()
	return int(counter.Load())
}

// CounterWithMutex функция для демонстрации решения состояния гонки
// при помощи примитива  Mutex из пакета sync
func CounterWithMutex(n int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex

	var counter int
	for range n {
		wg.Go(
			func() {
				mu.Lock()
				counter++
				mu.Unlock()
			},
		)
	}
	wg.Wait()
	return counter
}
