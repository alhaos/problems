package mergeChannels

import "sync"

// Условие задачи
// Нужно реализовать функцию, которая объединяет несколько входных каналов
// в один выходной канал, передавая в него все полученные значения
func merge(channels ...<-chan int) chan int {

	outCh := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				outCh <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}
