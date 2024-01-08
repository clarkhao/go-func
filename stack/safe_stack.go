package stack

import "sync"

// SafeStack is the thead safe stack
type SafeStack[T interface{}] struct {
	Stack Stacker[T]
	Lock  sync.RWMutex
}

// Len is the method that returns the size of stack
func (ss *SafeStack[T]) Len() int {
	ss.Lock.RLock()
	defer ss.Lock.RUnlock()
	return ss.Stack.Len()
}

// Push is the method that push the value into the stack
func (ss *SafeStack[T]) Push(v T, wg *sync.WaitGroup) {
	defer wg.Done()
	ss.Lock.Lock()
	defer ss.Lock.Unlock()
	ss.Stack.Push(v)
}

// Pop is the method that pop out from stack
// returns the value popped out, and indicate whether stack is empty
func (ss *SafeStack[T]) Pop(wg *sync.WaitGroup) (value T, ok bool) {
	defer wg.Done()
	ss.Lock.Lock()
	defer ss.Lock.Unlock()
	value, ok = ss.Stack.Pop()
	return value, ok
}

// Iterate is the method that write value into channel from the top one by one
func (ss *SafeStack[T]) Iterate() <-chan T {
	ss.Lock.Lock()
	defer ss.Lock.Unlock()
	return ss.Stack.Iterate()
}
