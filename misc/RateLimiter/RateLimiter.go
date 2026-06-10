package RateLimiter

import (
	"io"
	"net/http"
	"sync"
)

// Result структура содержит ответ
type Result struct {
	URL  string
	Body string
	Err  error
}

func Fetch(urls []string, maxConcurrent int) []Result {
	var wg sync.WaitGroup
	result := make([]Result, len(urls))
	semaphore := make(chan struct{}, maxConcurrent)
	for i := range len(urls) {
		result[i].URL = urls[i]
		wg.Add(1)
		go func(i int) {
			defer func() {
				<-semaphore
				wg.Done()
			}()
			semaphore <- struct{}{}
			resp, err := http.Get(urls[i])
			if err != nil {
				result[i].Err = err
				return
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				result[i].Err = err
				return
			}
			result[i].Body = string(body)
		}(i)
	}
	wg.Wait()
	return result
}
