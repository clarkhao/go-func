package stack

// Stacker is the interface including 4 methods
type Stacker[T interface{}] interface {
	Len() int
	Push(T)
	Pop() (value T, ok bool)
	Iterate() <-chan T
}

// Stack is the unsafe struct of stack
// inside the struct, element is a linkedlist
// size is the length of stack
type Stack[T interface{}] struct {
	top  *element[T]
	size int
}

type element[T interface{}] struct {
	value T
	next  *element[T]
}

// Len is the method that returns the size of stack
func (s *Stack[T]) Len() int {
	return s.size
}

// Push is the method that push the value into the stack
func (s *Stack[T]) Push(v T) {
	el := element[T]{
		value: v,
		next:  s.top,
	}
	s.top = &el
	s.size++
}

// Pop is the method that pop out from stack
// returns the value popped out, and indicate whether stack is empty
func (s *Stack[T]) Pop() (value T, ok bool) {
	if s.top == nil {
		return value, false
	}
	value = s.top.value
	s.top = s.top.next
	s.size--
	return value, true
}

// Iterate is the method that write value into channel from the top one by one
func (s *Stack[T]) Iterate() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		current := s.top
		for current != nil {
			ch <- current.value
			current = current.next
		}
	}()
	return ch
}
