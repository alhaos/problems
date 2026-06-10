# RateLimiter

## Задача

Реализовать конкурентный fetcher с ограничением максимального количества одновременных HTTP запросов.

## Сигнатура

```go
func Fetch(urls []string, maxConcurrent int) []Result
```

## Пример

```go
urls := []string{
    "https://example.com",
    "https://google.com",
    "https://github.com",
}

results := Fetch(urls, 2) // не более 2 запросов одновременно
for _, r := range results {
    if r.Err != nil {
        fmt.Println(r.URL, "error:", r.Err)
        continue
    }
    fmt.Println(r.URL, "body length:", len(r.Body))
}
```

## Требования

- Все запросы выполняются конкурентно
- Не более `maxConcurrent` запросов одновременно
- Порядок результатов соответствует порядку входных URL
- Ошибки не прерывают выполнение — сохраняются в `Result.Err`

## Как работает семафор

```go
semaphore := make(chan struct{}, maxConcurrent)

semaphore <- struct{}{}  // занять слот — блокируется если мест нет
// ... выполняем запрос ...
<-semaphore              // освободить слот
```

Буферизованный канал работает как турникет — пускает не более `maxConcurrent` горутин одновременно.
`struct{}` занимает 0 байт памяти — нам важен только факт занятости слота, не значение.

## Реализация 

```go

package RateLimiter

import (
	"io"
	"net/http"
	"sync"
)

// Result структура содержит ответ
type Result struct {
	// URL запроса
	URL string
	// Тело ответа
	Body string
	// Ошибка или nil
	Err error
}

// Fetch конкурентно возвращает содержимое urls в порядке слайса urls
// с ограничением максимального количества одновременных HTTP запросов.
func Fetch(urls []string, maxConcurrent int) []Result {

	// Инициализация WaitGroup
	var wg sync.WaitGroup

	// Инициализация слайса результатов длинной слайа urls
	result := make([]Result, len(urls))

	// Инициализация буферизированного канала, который выступает в качестве семафора,
	// при необходимости запустить горутину пишем туда пустую структуру (так как она не занимает места),
	// если буфер заполнен следующая горутина попытавшаяся записать в канал получит блокировку,
	// пока кто то не вычитает из канала, таким образом обеспечивается одновременная работа не более
	// maxConcurrent горутин
	semaphore := make(chan struct{}, maxConcurrent)

	// Инициализация цикла по слайсу urls
	for i := range len(urls) {

		// Присваиваем полю URL i-го результата значения i-го значения urls
		result[i].URL = urls[i]

		// Вызываем горутину (так не сработает в go < 1.25
		// несколько горутим могут обрабатывать один и то же i)
		wg.Go(
			func() {
				// В конце работы горутины вызываем процедуру в которой
				defer func() {
					// высвобождаем слот в семафоре
					<-semaphore
					// высвобождаем слот в WaitGroup
					wg.Done()
				}()
				// занимаем слот в семафоре, если удалось двигаемся дальше
				// если буфер полный сидим в блокировке
				semaphore <- struct{}{}

				// Выполняем запрос получаем resp и err
				resp, err := http.Get(urls[i])
				// Если была ошибка
				if err != nil {
					// У i-го результата заполняем поле Err
					result[i].Err = err
					// Завершаем горутину выполняется defer
					return
				}
				// Вычитываем тело из запроса
				body, err := io.ReadAll(resp.Body)
				// Если была ошибка
				if err != nil {
					// У i-го результата заполняем поле Err
					result[i].Err = err
					// Завершаем горутину выполняется defer
					return
				}
				// У i-го результата заполняем поле Body
				result[i].Body = string(body)
			},
		)
	}
	// Ждем пока все горутины завершат работу
	wg.Wait()
	// Возвращаем результат
	return result
}

```

## Тесты

```go

package RateLimiter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

// makeServer создаёт тестовый HTTP сервер с задержкой
func makeServer(t *testing.T, delay time.Duration) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		fmt.Fprintln(w, "ok")
	}))
}

func TestFetch_Basic(t *testing.T) {
	srv := makeServer(t, 0)
	defer srv.Close()

	urls := []string{srv.URL, srv.URL, srv.URL}
	results := Fetch(urls, 3)

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}
	for _, r := range results {
		if r.Err != nil {
			t.Errorf("unexpected error: %v", r.Err)
		}
		if r.Body == "" {
			t.Errorf("expected non-empty body for %s", r.URL)
		}
	}
}

func TestFetch_URLsPreserved(t *testing.T) {
	srv := makeServer(t, 0)
	defer srv.Close()

	urls := []string{srv.URL + "/a", srv.URL + "/b", srv.URL + "/c"}
	results := Fetch(urls, 3)

	for i, r := range results {
		if r.URL != urls[i] {
			t.Errorf("index %d: got URL %q, want %q", i, r.URL, urls[i])
		}
	}
}

func TestFetch_ErrorHandling(t *testing.T) {
	urls := []string{"http://localhost:1", "http://localhost:2"}
	results := Fetch(urls, 2)

	for _, r := range results {
		if r.Err == nil {
			t.Errorf("expected error for unreachable URL %s", r.URL)
		}
	}
}

func TestFetch_MaxConcurrent(t *testing.T) {
	var active atomic.Int32
	var maxSeen atomic.Int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cur := active.Add(1)
		defer active.Add(-1)

		// фиксируем максимум
		for {
			old := maxSeen.Load()
			if cur <= old || maxSeen.CompareAndSwap(old, cur) {
				break
			}
		}

		time.Sleep(20 * time.Millisecond)
		fmt.Fprintln(w, "ok")
	}))
	defer srv.Close()

	urls := make([]string, 10)
	for i := range urls {
		urls[i] = srv.URL
	}

	maxConcurrent := 3
	Fetch(urls, maxConcurrent)

	if got := int(maxSeen.Load()); got > maxConcurrent {
		t.Errorf("max concurrent requests = %d, want <= %d", got, maxConcurrent)
	}
}

func TestFetch_EmptyURLs(t *testing.T) {
	results := Fetch([]string{}, 3)
	if len(results) != 0 {
		t.Errorf("expected empty result, got %v", results)
	}
}

```

## Результаты

```text

PS C:\repo\problems\misc\RateLimiter> go test -v -race
=== RUN   TestFetch_Basic
--- PASS: TestFetch_Basic (0.20s)
=== RUN   TestFetch_URLsPreserved
--- PASS: TestFetch_URLsPreserved (0.01s)
=== RUN   TestFetch_ErrorHandling
--- PASS: TestFetch_ErrorHandling (0.03s)
=== RUN   TestFetch_MaxConcurrent
--- PASS: TestFetch_MaxConcurrent (0.10s)
=== RUN   TestFetch_EmptyURLs
--- PASS: TestFetch_EmptyURLs (0.00s)
PASS
ok  	github.com/alhaos/problems/misc/RateLimiter	1.741s

```

## Темы

`semaphore` `goroutine` `sync.WaitGroup` `http` `concurrency`
