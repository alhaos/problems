package towerOoHanoi

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

// Тест инициализации
func TestInit(t *testing.T) {
	var pegs Pegs
	n := 5
	pegs.Init(n)

	if pegs[pegA] != n {
		t.Errorf("Init() failed: pegA = %d, want %d", pegs[pegA], n)
	}
	if pegs[pegB] != 0 {
		t.Errorf("Init() failed: pegB = %d, want 0", pegs[pegB])
	}
	if pegs[pegC] != 0 {
		t.Errorf("Init() failed: pegC = %d, want 0", pegs[pegC])
	}
}

// Тест базового случая (n=1)
func TestSolveBaseCase(t *testing.T) {
	var pegs Pegs
	pegs.Init(1)

	// Перехватываем вывод
	output := captureOutput(func() {
		pegs.Solve(1, pegA, pegC, pegB)
	})

	expected := "Переместить диск со стержня A на стержень C\n"
	if output != expected {
		t.Errorf("Solve(1) output = %q, want %q", output, expected)
	}

	// Проверяем состояние стержней
	if pegs[pegA] != 0 {
		t.Errorf("After move: pegA = %d, want 0", pegs[pegA])
	}
	if pegs[pegC] != 1 {
		t.Errorf("After move: pegC = %d, want 1", pegs[pegC])
	}
}

// Тест для 2 дисков
func TestSolveFor2Disks(t *testing.T) {
	var pegs Pegs
	pegs.Init(2)

	output := captureOutput(func() {
		pegs.SolveTower(2)
	})

	// Ожидаемые перемещения для 2 дисков
	expectedMoves := []string{
		"Переместить диск со стержня A на стержень B",
		"Переместить диск со стержня A на стержень C",
		"Переместить диск со стержня B на стержень C",
	}

	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != len(expectedMoves) {
		t.Errorf("Expected %d moves, got %d", len(expectedMoves), len(lines))
	}

	for i, line := range lines {
		if line != expectedMoves[i] {
			t.Errorf("Move %d: got %q, want %q", i+1, line, expectedMoves[i])
		}
	}

	// Проверяем конечное состояние
	if pegs[pegA] != 0 || pegs[pegB] != 0 || pegs[pegC] != 2 {
		t.Errorf("Final state: A=%d, B=%d, C=%d, want A=0,B=0,C=2",
			pegs[pegA], pegs[pegB], pegs[pegC])
	}
}

// Тест для 3 дисков
func TestSolveFor3Disks(t *testing.T) {
	var pegs Pegs
	pegs.Init(3)

	output := captureOutput(func() {
		pegs.SolveTower(3)
	})

	// Количество перемещений для n дисков = 2^n - 1
	expectedMovesCount := 7
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != expectedMovesCount {
		t.Errorf("Expected %d moves, got %d", expectedMovesCount, len(lines))
	}

	// Проверяем конечное состояние
	if pegs[pegA] != 0 || pegs[pegB] != 0 || pegs[pegC] != 3 {
		t.Errorf("Final state: A=%d, B=%d, C=%d, want A=0,B=0,C=3",
			pegs[pegA], pegs[pegB], pegs[pegC])
	}
}

// Тест для 4 дисков
func TestSolveFor4Disks(t *testing.T) {
	var pegs Pegs
	pegs.Init(4)

	pegs.SolveTower(4)

	// Количество перемещений для 4 дисков = 15
	// Проверяем только конечное состояние для производительности
	if pegs[pegA] != 0 || pegs[pegB] != 0 || pegs[pegC] != 4 {
		t.Errorf("Final state: A=%d, B=%d, C=%d, want A=0,B=0,C=4",
			pegs[pegA], pegs[pegB], pegs[pegC])
	}
}

// Тест метода IsSolved
func TestIsSolved(t *testing.T) {
	var pegs Pegs
	n := 3

	pegs.Init(n)
	if pegs.IsSolved(n) {
		t.Error("IsSolved() returned true for unsolved puzzle")
	}

	pegs.SolveTower(n)
	if !pegs.IsSolved(n) {
		t.Error("IsSolved() returned false for solved puzzle")
	}
}

// Тест метода Print
func TestPrint(t *testing.T) {
	var pegs Pegs
	pegs.Init(5)
	pegs[pegB] = 2
	pegs[pegC] = 1

	output := captureOutput(func() {
		pegs.Print()
	})

	expected := "Стержень A: 5 дисков, B: 2 дисков, C: 1 дисков\n"
	if output != expected {
		t.Errorf("Print() = %q, want %q", output, expected)
	}
}

// Тест корректности количества перемещений (2^n - 1)
func TestMoveCount(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 3},
		{3, 7},
		{4, 15},
		{5, 31},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%d", tc.n), func(t *testing.T) {
			var pegs Pegs
			pegs.Init(tc.n)

			output := captureOutput(func() {
				pegs.SolveTower(tc.n)
			})

			moves := strings.Split(strings.TrimSpace(output), "\n")
			if len(moves) != tc.expected {
				t.Errorf("For n=%d: expected %d moves, got %d",
					tc.n, tc.expected, len(moves))
			}
		})
	}
}

// Тест, что все диски перемещены корректно (сумма дисков сохраняется)
func TestDiskConservation(t *testing.T) {
	testCases := []int{1, 2, 3, 4, 5}

	for _, n := range testCases {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			var pegs Pegs
			pegs.Init(n)

			// Сохраняем общее количество дисков
			totalBefore := pegs[pegA] + pegs[pegB] + pegs[pegC]

			pegs.SolveTower(n)

			totalAfter := pegs[pegA] + pegs[pegB] + pegs[pegC]

			if totalBefore != totalAfter {
				t.Errorf("Disk count mismatch: before=%d, after=%d",
					totalBefore, totalAfter)
			}
		})
	}
}

// Бенчмарк для измерения производительности
func BenchmarkSolveFor3Disks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var pegs Pegs
		pegs.Init(3)
		pegs.SolveTower(3)
	}
}

func BenchmarkSolveFor10Disks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var pegs Pegs
		pegs.Init(10)
		pegs.SolveTower(10)
	}
}

func BenchmarkSolveFor15Disks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var pegs Pegs
		pegs.Init(15)
		pegs.SolveTower(15)
	}
}

// Тест параллельного выполнения (если потребуется)
func TestConcurrentSolve(t *testing.T) {
	n := 3
	results := make(chan bool, 3)

	for range 3 {
		go func() {
			var pegs Pegs
			pegs.Init(n)
			pegs.SolveTower(n)
			results <- pegs.IsSolved(n)
		}()
	}

	for range 3 {
		if !<-results {
			t.Error("Concurrent solve failed")
		}
	}
}

// Вспомогательная функция для захвата вывода
func captureOutput(f func()) string {
	var buf bytes.Buffer
	// Сохраняем оригинальный вывод
	original := os.Stdout
	// Перенаправляем вывод в буфер
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Выполняем функцию
	f()

	// Восстанавливаем вывод
	w.Close()
	os.Stdout = original
	buf.ReadFrom(r)

	return buf.String()
}
