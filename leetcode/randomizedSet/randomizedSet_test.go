package randomizedSet

import (
	"testing"
)

func TestInsert(t *testing.T) {
	set := NewRandomizedSet()

	// Тест 1: Вставка нового элемента
	if !set.Insert(1) {
		t.Error("Insert(1) should return true")
	}

	// Тест 2: Вставка существующего элемента
	if set.Insert(1) {
		t.Error("Insert(1) should return false (already exists)")
	}

	// Тест 3: Вставка другого элемента
	if !set.Insert(2) {
		t.Error("Insert(2) should return true")
	}
}

func TestRemove(t *testing.T) {
	set := NewRandomizedSet()
	set.Insert(1)
	set.Insert(2)
	set.Insert(3)

	// Тест 1: Удаление существующего элемента
	if !set.Remove(2) {
		t.Error("Remove(2) should return true")
	}

	// Тест 2: Удаление несуществующего элемента
	if set.Remove(2) {
		t.Error("Remove(2) should return false (already removed)")
	}

	// Тест 3: Проверяем, что 1 все еще существует
	if !set.Remove(1) {
		t.Error("Remove(1) should return true (still exists)")
	}

	// Тест 4: Проверяем, что 3 все еще существует
	if !set.Remove(3) {
		t.Error("Remove(3) should return true (still exists)")
	}

	// Тест 5: Проверяем, что множество пусто
	if set.Remove(1) {
		t.Error("Remove(1) should return false (already removed)")
	}
	if set.Remove(3) {
		t.Error("Remove(3) should return false (already removed)")
	}
}

func TestGetRandom(t *testing.T) {
	set := NewRandomizedSet()
	set.Insert(1)
	set.Insert(2)
	set.Insert(3)

	// Тест: GetRandom должен возвращать только существующие элементы
	for _ = range 100 {
		val := set.GetRandom()
		if val != 1 && val != 2 && val != 3 {
			t.Errorf("GetRandom() returned %d, expected 1, 2 or 3", val)
		}
	}
}

func TestComplexScenario(t *testing.T) {
	set := NewRandomizedSet()

	// Сценарий 1: Вставка и удаление
	set.Insert(1)
	set.Insert(2)
	set.Insert(3)
	set.Remove(2)

	// Проверяем, что 2 нет
	if set.Remove(2) {
		t.Error("Remove(2) should return false after removal")
	}

	// Сценарий 2: Удаление последнего элемента
	set.Insert(4)
	set.Remove(4)
	if set.Remove(4) {
		t.Error("Remove(4) should return false after removal")
	}

	// Сценарий 3: Вставка после удаления
	set.Insert(5)
	if !set.Insert(6) {
		t.Error("Insert(6) should return true")
	}
}

func TestRandomizedSet_RemoveEdgeCases(t *testing.T) {
	set := NewRandomizedSet()
	set.Insert(1)
	set.Insert(2)
	set.Insert(3)
	set.Insert(4)

	// Удаляем последний элемент
	if !set.Remove(4) {
		t.Error("Remove(4) should return true")
	}

	// Удаляем первый элемент
	if !set.Remove(1) {
		t.Error("Remove(1) should return true")
	}

	// Проверяем оставшиеся элементы
	if set.Remove(1) {
		t.Error("Remove(1) should return false")
	}
	if set.Remove(4) {
		t.Error("Remove(4) should return false")
	}
}

func TestRandomizedSet_StressTest(t *testing.T) {
	set := NewRandomizedSet()
	n := 1000

	// Вставляем n элементов
	for i := 0; i < n; i++ {
		if !set.Insert(i) {
			t.Errorf("Insert(%d) should return true", i)
		}
	}

	// Проверяем, что все элементы вставлены
	for i := 0; i < n; i++ {
		if set.Insert(i) {
			t.Errorf("Insert(%d) should return false (already exists)", i)
		}
	}

	// Удаляем каждый второй элемент (четные)
	for i := 0; i < n; i += 2 {
		if !set.Remove(i) {
			t.Errorf("Remove(%d) should return true", i)
		}
	}

	// Проверяем, что четные удалены, нечетные остались
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			// Нечетные должны существовать
			if !set.Remove(i) {
				t.Errorf("Remove(%d) should return true (should still exist)", i)
			}
		} else {
			// Четные должны быть удалены
			if set.Remove(i) {
				t.Errorf("Remove(%d) should return false (already removed)", i)
			}
		}
	}
}

func TestRandomizedSet_MultipleOperations(t *testing.T) {
	set := NewRandomizedSet()

	// Последовательность операций
	operations := []struct {
		op   string
		val  int
		want bool
	}{
		{"insert", 1, true},
		{"insert", 2, true},
		{"insert", 3, true},
		{"insert", 1, false},
		{"remove", 2, true},
		{"remove", 2, false},
		{"insert", 4, true},
		{"remove", 1, true},
		{"remove", 3, true},
		{"remove", 4, true},
	}

	for _, op := range operations {
		var got bool
		if op.op == "insert" {
			got = set.Insert(op.val)
		} else {
			got = set.Remove(op.val)
		}

		if got != op.want {
			t.Errorf("%s(%d) = %v, want %v", op.op, op.val, got, op.want)
		}
	}

	// Проверяем, что все удалено
	if len(set.values) != 0 || len(set.data) != 0 {
		t.Errorf("Set should be empty, values=%v, data=%v", set.values, set.data)
	}
}
