package queue

// Queuer is the interface including 5 methods
type Queuer[T interface{}] interface {
	Len() int
	Push(v T)
	Pop() (value T, ok bool)
	IterateFromHead() <-chan T
	IterateFromTail() <-chan T
}

// UnsafeQueue is the struct for queue
// made up of double linked list and size(int)
type UnsafeQueue[T interface{}] struct {
	head *element[T]
	tail *element[T]
	size int
}

type element[T interface{}] struct {
	value T
	next  *element[T]
	prev  *element[T]
}

// Len is the method that returns the length of queue
func (q *UnsafeQueue[T]) Len() int {
	return q.size
}

// Push is the method that push item v into queue
func (q *UnsafeQueue[T]) Push(v T) {
	el := &element[T]{
		value: v,
		next:  q.tail,
		prev:  nil,
	}
	switch q.size {
	case 0:
		q.head = el
		q.tail = el
	case 1:
		q.head.prev = el
		q.tail = el
	default:
		q.tail.prev = el
		q.tail = el
	}
	q.size++
}

// Pop is the method that pop out value and indicate whether queue is empty
func (q *UnsafeQueue[T]) Pop() (value T, ok bool) {
	if q.head == nil {
		return value, false
	}
	value = q.head.value
	switch q.size {
	case 1:
		q.head = q.head.prev
		q.tail = q.tail.next
	default:
		q.head = q.head.prev
		q.head.next = nil
	}
	q.size--
	return value, true
}

// IterateFromHead is the method that itearate value of queue and write into chan one by one from the head
func (q *UnsafeQueue[T]) IterateFromHead() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		current := q.head
		for current != nil {
			ch <- current.value
			current = current.prev
		}
	}()
	return ch
}

// IterateFromHead is the method that itearate value of queue and write into chan one by one from the tail
func (q *UnsafeQueue[T]) IterateFromTail() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		current := q.tail
		for current != nil {
			ch <- current.value
			current = current.next
		}
	}()
	return ch
}
