// Package pipeline реализует паттерн конвейерной обработки данных (pipeline).
//
// Pipeline — это цепочка горутин, где каждая стадия получает данные
// из входного канала, обрабатывает их и передаёт в выходной канал:
//
//	[generate] → chan → [toUpper] → chan → [filterPrefix] → chan
//
// Каждая стадия:
//   - запускает собственную горутину
//   - закрывает выходной канал через defer когда данные закончились
//   - реагирует на отмену через context.Context
package pipeline

import (
	"context"
	"strings"
)

// Pipeline возвращает канал string который выдает последовательно
// обработанные слова из words
func Pipeline(ctx context.Context, prefix string, words ...string) <-chan string {

	// Первый элемент конвейера функция generate
	gen := generate(ctx, words...)
	up := toUpper(ctx, gen)
	return filterPrefix(ctx, up, prefix)
}

// generate последовательно выводит в возвращаемый канал полученные слова.
// Завершает выполнение при завершении контекста.
func generate(ctx context.Context, words ...string) <-chan string {

	// Инициируем канал
	out := make(chan string)

	// Вызываем горутину которая записывает в канал слова если их кто то читаем
	// и завершает выполнение в случае если контекст завершился предварительно
	go func() {
		// Закрываем канал после выполнения горутины
		defer close(out)
		// Цикл range по полученным words
		for _, word := range words {
			// Конструкция select блокирует выполнение до момента пока не произойдет
			// одно из возможных событий
			select {
			// На стороне канала out кто то готов прочитать слово
			case out <- word:
			// Контекст выполнения завершен.
			case <-ctx.Done():
				// Останавливаем выполнение
				return
			}
		}
	}()
	// Возвращаем канал
	return out
}

// toUpper второй элемент конвейера получает слова из канала
// переводит их в верхний регистр и отправляет в дальше по конвейеру в новый канал
// Завершает выполнение при завершении контекста.
func toUpper(ctx context.Context, in <-chan string) <-chan string {
	// Инициируем канал
	out := make(chan string)
	// Вызываем горутину
	go func() {
		// Закрываем канал после выполнения горутины
		defer close(out)
		// Читаем из канала используя цикл range
		for word := range in {
			// Конструкция select блокирует выполнение до момента пока не произойдет
			// одно из возможных событий
			select {
			// На стороне канала out кто то готов прочитать слово
			case out <- strings.ToUpper(word):
			// Контекст выполнения завершен.
			case <-ctx.Done():
				// Останавливаем выполнение
				return
			}
		}
	}()
	// Возвращаем канал
	return out
}

// filterPrefix третий элемент конвейера получает слова из канала
// в случае если слово имеет префикс prefix передает его далее по конвейеру
// Завершает выполнение при завершении контекста.
func filterPrefix(ctx context.Context, in <-chan string, prefix string) <-chan string {
	// Инициируем канал
	out := make(chan string)
	// Вызываем горутину
	go func() {
		// Закрываем канал после выполнения горутины
		defer close(out)
		// Читаем из канала используя цикл range
		for word := range in {
			// Проверяем наличие префикса
			if strings.HasPrefix(word, prefix) {
				// Конструкция select блокирует выполнение до момента пока не произойдет
				// одно из возможных событий
				select {
				// На стороне канала out кто то готов прочитать слово
				case out <- word:
				// Контекст выполнения завершен.
				case <-ctx.Done():
					// Останавливаем выполнение
					return
				}
			}
		}
	}()
	// Возвращаем канал
	return out
}
