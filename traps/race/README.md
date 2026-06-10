# Состояние гонки (race condition)

Состояние гонки (Race Condition) — это ошибка, возникающая при одновременном
доступе нескольких горутин к одним и тем же данным без должной синхронизации.
При этом результат выполнения программы начинает зависеть от случайного порядка
выполнения потоков, что приводит к непредсказуемым результатам.

## Модель

### Функция

```go

package race

import (
  "sync"
)

// CounterWithRace функция для моделирования состояния гонки
// принимает на вход n - количество выполняемых итераций цикла
func CounterWithRace(n int) int {
	
	// Инициируем WaitGroup чтобы главная горутина не завершилась раньше вложенных
	// горутин
	var wg sync.WaitGroup
	
	// Инициируем переменную counter в которую будем инкрементировать в каждом
	// цикле 
	var counter int
	
	// Запускаем цикл с количеством итераций n
	for range n {
		// wg.Go занимает место в в WaitGroup запускает функцию аргумент отдельной
		// горутиной и высвобождает при выполнении функции, работает с 1.25
		wg.Go(
			func() {
				// Вот тут происходит само действо
				// counter++ — это не атомарная операция, а три шага
				// 1. READ  — загрузить значение counter из памяти в регистр
				// 2. ADD   — прибавить 1 в регистре
				// 3. WRITE — записать результат обратно в память
				// 
				// Горутина A          Горутина B
				// READ  → 5
        //                     READ  → 5      ← читает ДО того как A записала
        // ADD   → 6
        //                     ADD   → 6
        // WRITE → counter=6
        //                     WRITE → counter=6   ← затирает результат A!
        // Итог: две горутины сделали ++, а счётчик вырос только на 1. 
        // Это lost update.
				counter++
			},
		)
	}
	// WaitGroup дожидается завершения всех горутин
	wg.Wait()
	// Возвращаем counter
	return counter
}

```

### Тест

```go

package race

import "testing"

func TestCounterWithRace(t *testing.T) {
	testCases := []struct {
		desc     string
		n        int
		expected int
	}{
		{
			desc:     "1 итерация",
			n:        1,
			expected: 1,
		},
		{
			desc:     "10 итераций",
			n:        10,
			expected: 10,
		},
		{
			desc:     "100 итераций",
			n:        100,
			expected: 100,
		},
		{
			desc:     "1000 итераций",
			n:        1000,
			expected: 1000,
		},
		{
			desc:     "10000 итераций",
			n:        10000,
			expected: 10000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := CounterWithRace(tC.n)
			if result != tC.expected {
				t.Errorf("unexpected result for test %s, expected: %d, but got: %d", tC.desc, tC.expected, result)
			}
		})
	}
}

```

### Результат

Самое неприятное, что заметить это мы можем часто только на достаточно большом количестве повторений.

```text
PS C:\repo\problems\traps\race> go test -v
=== RUN   TestCounterWithRace
=== RUN   TestCounterWithRace/1_итерация
=== RUN   TestCounterWithRace/10_итераций
=== RUN   TestCounterWithRace/100_итераций
=== RUN   TestCounterWithRace/1000_итераций
=== RUN   TestCounterWithRace/10000_итераций
    race_test.go:41: unexpected result for test 10000 итераций, expected: 10000, but got: 9408
--- FAIL: TestCounterWithRace (0.02s)
    --- PASS: TestCounterWithRace/1_итерация (0.00s)
    --- PASS: TestCounterWithRace/10_итераций (0.00s)
    --- PASS: TestCounterWithRace/100_итераций (0.00s)
    --- PASS: TestCounterWithRace/1000_итераций (0.01s)
    --- FAIL: TestCounterWithRace/10000_итераций (0.01s)
FAIL
exit status 1
FAIL	github.com/alhaos/problems/traps/race	0.238s

```

Тест на 10000 итераций упал с неверным значением, а меньшее количество итераций проблему не выявило.

В тестах го есть встроенный механизм детектирования состояния гонки.

```text

PS C:\repo\problems\traps\race> go test -race
==================
WARNING: DATA RACE
Read at 0x00c00000c0f8 by goroutine 20:
  github.com/alhaos/problems/traps/race.CounterWithRace.func1()
      C:/repo/problems/traps/race/race.go:13 +0x2e
  sync.(*WaitGroup).Go.func1()
      C:/Program Files/Go/src/sync/waitgroup.go:258 +0x5d

Previous write at 0x00c00000c0f8 by goroutine 12:
  github.com/alhaos/problems/traps/race.CounterWithRace.func1()
      C:/repo/problems/traps/race/race.go:13 +0x44
  sync.(*WaitGroup).Go.func1()
      C:/Program Files/Go/src/sync/waitgroup.go:258 +0x5d

Goroutine 20 (running) created at:
  sync.(*WaitGroup).Go()
      C:/Program Files/Go/src/sync/waitgroup.go:238 +0x86
  github.com/alhaos/problems/traps/race.CounterWithRace()
      C:/repo/problems/traps/race/race.go:11 +0x8f
  github.com/alhaos/problems/traps/race.TestCounterWithRace.func1()
      C:/repo/problems/traps/race/race_test.go:39 +0x64
  testing.tRunner()
      C:/Program Files/Go/src/testing/testing.go:2036 +0x1ca
  testing.(*T).Run.gowrap1()
      C:/Program Files/Go/src/testing/testing.go:2101 +0x38

Goroutine 12 (finished) created at:
  sync.(*WaitGroup).Go()
      C:/Program Files/Go/src/sync/waitgroup.go:238 +0x86
  github.com/alhaos/problems/traps/race.CounterWithRace()
      C:/repo/problems/traps/race/race.go:11 +0x8f
  github.com/alhaos/problems/traps/race.TestCounterWithRace.func1()
      C:/repo/problems/traps/race/race_test.go:39 +0x64
  testing.tRunner()
      C:/Program Files/Go/src/testing/testing.go:2036 +0x1ca
  testing.(*T).Run.gowrap1()
      C:/Program Files/Go/src/testing/testing.go:2101 +0x38
==================
==================
WARNING: DATA RACE
Write at 0x00c000208088 by goroutine 137:
  github.com/alhaos/problems/traps/race.CounterWithRace.func1()
      C:/repo/problems/traps/race/race.go:13 +0x44
  sync.(*WaitGroup).Go.func1()
      C:/Program Files/Go/src/sync/waitgroup.go:258 +0x5d

Previous write at 0x00c000208088 by goroutine 139:
  github.com/alhaos/problems/traps/race.CounterWithRace.func1()
      C:/repo/problems/traps/race/race.go:13 +0x44
  sync.(*WaitGroup).Go.func1()
      C:/Program Files/Go/src/sync/waitgroup.go:258 +0x5d

Goroutine 137 (running) created at:
  sync.(*WaitGroup).Go()
      C:/Program Files/Go/src/sync/waitgroup.go:238 +0x86
  github.com/alhaos/problems/traps/race.CounterWithRace()
      C:/repo/problems/traps/race/race.go:11 +0x8f
  github.com/alhaos/problems/traps/race.TestCounterWithRace.func1()
      C:/repo/problems/traps/race/race_test.go:39 +0x64
  testing.tRunner()
      C:/Program Files/Go/src/testing/testing.go:2036 +0x1ca
  testing.(*T).Run.gowrap1()
      C:/Program Files/Go/src/testing/testing.go:2101 +0x38

Goroutine 139 (finished) created at:
  sync.(*WaitGroup).Go()
      C:/Program Files/Go/src/sync/waitgroup.go:238 +0x86
  github.com/alhaos/problems/traps/race.CounterWithRace()
      C:/repo/problems/traps/race/race.go:11 +0x8f
  github.com/alhaos/problems/traps/race.TestCounterWithRace.func1()
      C:/repo/problems/traps/race/race_test.go:39 +0x64
  testing.tRunner()
      C:/Program Files/Go/src/testing/testing.go:2036 +0x1ca
  testing.(*T).Run.gowrap1()
      C:/Program Files/Go/src/testing/testing.go:2101 +0x38
==================
--- FAIL: TestCounterWithRace (0.10s)
    --- FAIL: TestCounterWithRace/10_итераций (0.00s)
        testing.go:1712: race detected during execution of test
    --- FAIL: TestCounterWithRace/1000_итераций (0.01s)
        testing.go:1712: race detected during execution of test
    --- FAIL: TestCounterWithRace/10000_итераций (0.08s)
        race_test.go:41: unexpected result for test 10000 итераций, expected: 10000, but got: 9768
FAIL
exit status 1
FAIL	github.com/alhaos/problems/traps/race	0.550s

```

## Решение на основе примитивов пакета atomic

Один из вариантов решить эту проблему использовать специальные потокобезопасные структуры из пакета atomic 

### Функция

```go

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

```

### Тест

```go

func TestCounterWithAtomic(t *testing.T) {
	testCases := []struct {
		desc     string
		n        int
		expected int
	}{
		{
			desc:     "1 итерация",
			n:        1,
			expected: 1,
		},
		{
			desc:     "10 итераций",
			n:        10,
			expected: 10,
		},
		{
			desc:     "100 итераций",
			n:        100,
			expected: 100,
		},
		{
			desc:     "1000 итераций",
			n:        1000,
			expected: 1000,
		},
		{
			desc:     "10000 итераций",
			n:        10000,
			expected: 10000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := CounterWithAtomic(tC.n)
			if result != tC.expected {
				t.Errorf("unexpected result for test %s, expected: %d, but got: %d", tC.desc, tC.expected, result)
			}
		})
	}
}

```

### Результат

```text

PS C:\repo\problems\traps\race> go test -v -race -run TestCounterWithAtomic
=== RUN   TestCounterWithAtomic
=== RUN   TestCounterWithAtomic/1_итерация
=== RUN   TestCounterWithAtomic/10_итераций
=== RUN   TestCounterWithAtomic/100_итераций
=== RUN   TestCounterWithAtomic/1000_итераций
=== RUN   TestCounterWithAtomic/10000_итераций
--- PASS: TestCounterWithAtomic (0.07s)
    --- PASS: TestCounterWithAtomic/1_итерация (0.00s)
    --- PASS: TestCounterWithAtomic/10_итераций (0.00s)
    --- PASS: TestCounterWithAtomic/100_итераций (0.00s)
    --- PASS: TestCounterWithAtomic/1000_итераций (0.00s)
    --- PASS: TestCounterWithAtomic/10000_итераций (0.06s)
PASS
ok  	github.com/alhaos/problems/traps/race	1.346s
PS C:\repo\problems\traps\race>

```

## Решение на основе примитива Mutex пакета sync

Еще один вариант - использовать Mutex.
Mutex предпочтительнее atomic: когда нужно защитить не одну переменную, а блок операций или составную структуру.

### Функция

```go

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

```

### Тест

```go

func TestCounterWithMutex(t *testing.T) {
	testCases := []struct {
		desc     string
		n        int
		expected int
	}{
		{
			desc:     "1 итерация",
			n:        1,
			expected: 1,
		},
		{
			desc:     "10 итераций",
			n:        10,
			expected: 10,
		},
		{
			desc:     "100 итераций",
			n:        100,
			expected: 100,
		},
		{
			desc:     "1000 итераций",
			n:        1000,
			expected: 1000,
		},
		{
			desc:     "10000 итераций",
			n:        10000,
			expected: 10000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := CounterWithMutex(tC.n)
			if result != tC.expected {
				t.Errorf("unexpected result for test %s, expected: %d, but got: %d", tC.desc, tC.expected, result)
			}
		})
	}
}

```

### Результат

```text

PS C:\repo\problems\traps\race> go test -v -race -run TestCounterWithMutex
=== RUN   TestCounterWithMutex
=== RUN   TestCounterWithMutex/1_итерация
=== RUN   TestCounterWithMutex/10_итераций
=== RUN   TestCounterWithMutex/100_итераций
=== RUN   TestCounterWithMutex/1000_итераций
=== RUN   TestCounterWithMutex/10000_итераций
--- PASS: TestCounterWithMutex (0.08s)
    --- PASS: TestCounterWithMutex/1_итерация (0.00s)
    --- PASS: TestCounterWithMutex/10_итераций (0.00s)
    --- PASS: TestCounterWithMutex/100_итераций (0.00s)
    --- PASS: TestCounterWithMutex/1000_итераций (0.01s)
    --- PASS: TestCounterWithMutex/10000_итераций (0.07s)
PASS
ok  	github.com/alhaos/problems/traps/race	1.415s

```
