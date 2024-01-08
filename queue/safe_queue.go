package queue

import "sync"

// Queue is the thread safe queue
type Queue[T interface{}] struct {
	Queue Queuer[T]
	Lock  sync.RWMutex
}

// Len is the method that returns the length of queue
func (q *Queue[T]) Len() int {
	q.Lock.RLock()
	defer q.Lock.RUnlock()
	return q.Queue.Len()
}

// Push is the method that push item v into queue
func (q *Queue[T]) Push(v T, wg *sync.WaitGroup) {
	defer wg.Done()
	q.Lock.Lock()
	defer q.Lock.Unlock()
	q.Queue.Push(v)
}

// Pop is the method that pop out value and indicate whether queue is empty
func (q *Queue[T]) Pop(wg *sync.WaitGroup) (value T, ok bool) {
	defer wg.Done()
	q.Lock.Lock()
	defer q.Lock.Unlock()
	return q.Queue.Pop()
}

// IterateFromHead is the method that itearate value of queue and write into chan one by one from the head
func (q *Queue[T]) IterateFromHead() <-chan T {
	q.Lock.Lock()
	defer q.Lock.Unlock()
	return q.Queue.IterateFromHead()
}

// IterateFromHead is the method that itearate value of queue and write into chan one by one from the tail
func (q *Queue[T]) IterateFromTail() <-chan T {
	q.Lock.Lock()
	defer q.Lock.Unlock()
	return q.Queue.IterateFromTail()
}
