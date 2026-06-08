# Pipeline

## Задача

Реализовать pipeline обработки строк из трёх последовательных стадий с поддерж-
кой отмены через `context.Context`.

## Сигнатуры

```go
// Точка входа — собирает все стадии в единый pipeline
func Pipeline(ctx context.Context, prefix string, words ...string) <-chan string

// Стадия 1: подаёт строки в канал
func Generate(ctx context.Context, words ...string) <-chan string

// Стадия 2: переводит каждую строку в верхний регистр
func ToUpper(ctx context.Context, in <-chan string) <-chan string

// Стадия 3: пропускает только строки с заданным префиксом
func FilterPrefix(ctx context.Context, in <-chan string, prefix string) <-chan string
```

## Пример

```go
ctx := context.Background()
out := Pipeline(ctx, "HE", "hello", "world", "help", "go")
for s := range out {
    fmt.Println(s)
}
// Output:
// HELLO
// HELP
```

## Требования

- Каждая стадия запускает собственную горутину
- Каждая горутина завершается при закрытии входного канала **или** при отмене контекста
- Выходной канал закрывается через `defer close(out)` — получатель узнаёт о
- завершении через `range`
- Утечек горутин быть не должно

## Паттерн каждой стадии

```go

func Stage(ctx context.Context, in <-chan T) <-chan T {
    out := make(chan T)
    go func() {
        defer close(out)
        for v := range in {
            select {
            case out <- process(v):
            case <-ctx.Done():
                return
            }
        }
    }()
    return out
}

```

## Темы

`pipeline` `goroutine` `channel` `context` `concurrency`
