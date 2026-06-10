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

Буферизованный канал работает как турникет — пускает не более `maxConcurrent` горутин одновременно. `struct{}` занимает 0 байт памяти — нам важен только факт занятости слота, не значение.

## Темы

`semaphore` `goroutine` `sync.WaitGroup` `http` `concurrency`
