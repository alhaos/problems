package priorityqueue

import (
	"context"
	"errors"
	"sync"
)

type Message struct {
	PartitionKey string      // Ключ для шардирования
	Priority     int         // Приоритет (чем меньше, тем выше)
	Payload      interface{} // Полезная нагрузка
}

var (
	ErrQueueClosed = errors.New("queue is closed")
	ErrQueueFull   = errors.New("queue is full")
	ErrContextDone = errors.New("context cancelled")
)

type Queue interface {

	// Push добавляет сообщение в очередь.
	// Если очередь заполнена, должна блокироваться до появления места или отмены контекста.
	// Возвращает ошибку, если ctx отменён до принятия сообщения.
	Push(ctx context.Context, msg Message) error

	// Pop возвращает следующее доступное сообщение для воркера.
	// Очередь гарантирует, что для одного PartitionKey сообщения возвращаются
	// в точном порядке их добавления.
	// Блокируется, если нет сообщений для этого ключа или все ключи сейчас заняты.
	Pop(ctx context.Context) (Message, error)

	// Close инициирует graceful shutdown.
	// Останавливает приём новых сообщений (Push будет возвращать ошибку) и
	// ожидает завершения обработки всех ранее полученных сообщений.
	Close() error

	// Ack уведомляет очередь о завершении обработки сообщения.
	// Позволяет следующему сообщению с тем же PartitionKey быть выданным.
	Ack(msg Message)
}

// queue структура обеспечивающая шардированную приоритетную очередь с гарантией порядка
type queue struct {
	// mu мютекс
	mu sync.Mutex
	// shards шарды сообщений
	shards map[string][]Message
	// locked карта флагов шардов находящихся в обработке
	locked map[string]bool
	// capacity емкость очереди
	capacity int
	// closed флаг признака закрытия очереди
	closed bool
	// notEmpty флаг признака пустой очереди
	notEmpty *sync.Cond
	// notFull флаг признака полной очереди
	notFull *sync.Cond
}

// NewQueue конструктор структуры queue
func NewQueue() Queue {

	// Инициируем карты
	q := &queue{
		shards: make(map[string][]Message),
		locked: make(map[string]bool),
	}

	// Инициируем Cord
	q.notEmpty = sync.NewCond(&q.mu)
	q.notFull = sync.NewCond(&q.mu)
	return q
}

// Добавляет сообщение в очередь
func (q *queue) Push(ctx context.Context, msg Message) error {

	// Устанавливаем блокировку
	q.mu.Lock()

	// Отложено снимаем блокировку при любом исходе выполнения функции
	defer q.mu.Unlock()

	// Проверяем закрыта ли очередь
	if q.closed {
		// Если да возвращаем ошибку
		return ErrQueueClosed
	}

	// Создаем канал чтобы прочитать из него когда он будет
	// закрыт для проверки что Push отработал
	stopWatch := make(chan struct{})

	// Отложено закрываем канал
	defer close(stopWatch)

	// Запускаем горутину для проверки отмены по контексту
	go func() {
		// Ожидаем момента пока не случится одно из событий
		select {
		// Вычитали из канала завершения контекста
		case <-ctx.Done():
			// Берем блокировку нужна для q.notFull.Broadcast()
			// не вызывает дедлок так как в отдельной горутине
			q.mu.Lock()
			// Будем всех кто ждет q.notFull,Wait
			q.notFull.Broadcast()
			// Снимаем блокировку
			q.mu.Unlock()
		case <-stopWatch:
			// Родительская горутина отработала канал закрылся
			// через defer close() и чтение из закрытого канала
			// немедленно возвращает дефолтное значение оно нам
			// не интересно, интересен факт чтения из канала
		}
	}()

	// Проверить заполнена ли очередь
	for q.totalLen() >= q.capacity {
		// если да
		// ждать notFull.Wait() в цикле.
		// горутина засыпает не жрёт CPU, пока кто-то не освободит место.
		q.notFull.Wait()
		// проверяем не отменен ли контекст
		if ctx.Err() != nil {
			return ErrContextDone
		}
	}

	// Добавить сообщение в shards[msg.PartitionKey]
	q.shards[msg.PartitionKey] = append(q.shards[msg.PartitionKey], msg)

	// Разбудить ожидающий Pop через
	q.notEmpty.Signal()

	return nil
}

func (q *queue) Pop(ctx context.Context) (Message, error) {
	panic("implement me")
}

func (q *queue) Close() error {
	panic("implement me")
}

func (q *queue) Ack(msg Message) {
	panic("implement me")
}

// totalLen вычисляет общую длину очереди
func (q *queue) totalLen() int {
	total := 0
	for _, msgs := range q.shards {
		total += len(msgs)
	}
	return total
}
