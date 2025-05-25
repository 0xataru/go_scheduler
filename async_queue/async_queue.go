package async_queue

import "sync"

type Queue[T any] struct {
	mx sync.RWMutex
	q  []T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		q: make([]T, 0),
	}
}

func (qu *Queue[T]) Put(item T) {
	qu.mx.Lock()
	qu.q = append(qu.q, item)
	qu.mx.Unlock()
}

func (qu *Queue[T]) Flush() []T {
	qu.mx.Lock()
	defer qu.mx.Unlock()
	res := qu.q
	qu.q = make([]T, 0)
	return res
}

func (qu *Queue[T]) Remove(predicate func(T) bool) {
	qu.mx.Lock()
	defer qu.mx.Unlock()

	var newQueue []T
	for _, item := range qu.q {
		if !predicate(item) {
			newQueue = append(newQueue, item)
		}
	}
	qu.q = newQueue
}
