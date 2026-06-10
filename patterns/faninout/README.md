# Fan-out / Fan-in

## Задача

Реализовать два паттерна конкурентной обработки данных и объединить их в единую
функцию.

## Сигнатуры

```go
// FanOutIn — объединяет Fan-out и Fan-in в единый pipeline
func FanOutIn(ctx context.Context, numWorkers int, in <-chan int) <-chan int

// FanOut — раздаёт данные из одного канала по numWorkers каналам (round-robin)
func FanOut(ctx context.Context, numWorkers int, in <-chan int) []<-chan int

// FanIn — сливает несколько каналов в один
func FanIn(ctx context.Context, channels ...<-chan int) <-chan int
```

## Пример

```go
ctx := context.Background()

in := make(chan int, 5)
for _, n := range []int{1, 2, 3, 4, 5} {
    in <- n
}
close(in)

out := FanOutIn(ctx, 3, in)
for r := range out {
    fmt.Println(r) // 1, 4, 9, 16, 25 (порядок не гарантирован)
}
```

## Как это работает

```text
         ┌──► ch1 ──► [worker1: x*x] ─┐
in ──────┼──► ch2 ──► [worker2: x*x] ─┼──► out
         └──► ch3 ──► [worker3: x*x] ─┘
            FanOut                   FanIn
```

**FanOut** — горутина-диспетчер читает из `in` и раздаёт значения по каналам
циклически (round-robin). Когда `in` закрыт или контекст отменён — закрывает
все выходные каналы.

**FanIn** — на каждый входной канал запускается горутина которая читает и пишет
в общий `out`. `sync.WaitGroup` следит за завершением всех горутин — только
тогда `out` закрывается.

## Отличие от Worker Pool

| Worker Pool                    | Fan-out / Fan-in                      |
|--------------------------------|---------------------------------------|
| один общий канал задач         | отдельный канал на каждый воркер      |
| Go runtime распределяет задачи | диспетчер явно раздаёт по round-robin |
| проще в реализации             | явный контроль над распределением     |

## Темы

`fan-out` `fan-in` `goroutine` `channel` `sync.WaitGroup` `context` `concurrency`
