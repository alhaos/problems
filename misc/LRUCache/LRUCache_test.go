package LRUCache

import "testing"

func TestLRUCache(t *testing.T) {
	t.Run("базовый пример из условия", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		assertGet(t, cache, 1, 1)
		cache.Put(3, 3) // вытесняем 2
		assertGet(t, cache, 2, -1)
		assertGet(t, cache, 3, 3)
	})

	t.Run("вытеснение самого старого при переполнении", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(3, 3) // вытесняем 1
		assertGet(t, cache, 1, -1)
		assertGet(t, cache, 2, 2)
		assertGet(t, cache, 3, 3)
	})

	t.Run("get обновляет порядок использования", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		assertGet(t, cache, 1, 1) // 1 становится свежим
		cache.Put(3, 3)           // вытесняем 2, не 1
		assertGet(t, cache, 1, 1)
		assertGet(t, cache, 2, -1)
		assertGet(t, cache, 3, 3)
	})

	t.Run("обновление существующего ключа", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(1, 10) // обновляем значение ключа 1
		assertGet(t, cache, 1, 10)
		assertGet(t, cache, 2, 2)
	})

	t.Run("обновление ключа обновляет порядок", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(1, 10) // 1 становится свежим
		cache.Put(3, 3)  // вытесняем 2, не 1
		assertGet(t, cache, 1, 10)
		assertGet(t, cache, 2, -1)
		assertGet(t, cache, 3, 3)
	})

	t.Run("кэш ёмкостью 1", func(t *testing.T) {
		cache := Constructor(1)
		cache.Put(1, 1)
		assertGet(t, cache, 1, 1)
		cache.Put(2, 2) // вытесняем 1
		assertGet(t, cache, 1, -1)
		assertGet(t, cache, 2, 2)
	})

	t.Run("get несуществующего ключа", func(t *testing.T) {
		cache := Constructor(3)
		assertGet(t, cache, 99, -1)
	})

	t.Run("несколько вытеснений подряд", func(t *testing.T) {
		cache := Constructor(3)
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(3, 3)
		cache.Put(4, 4) // вытесняем 1
		cache.Put(5, 5) // вытесняем 2
		assertGet(t, cache, 1, -1)
		assertGet(t, cache, 2, -1)
		assertGet(t, cache, 3, 3)
		assertGet(t, cache, 4, 4)
		assertGet(t, cache, 5, 5)
	})
}

func assertGet(t *testing.T, cache LRUCache, key, want int) {
	t.Helper()
	got := cache.Get(key)
	if got != want {
		t.Errorf("Get(%d) = %d, want %d", key, got, want)
	}
}
