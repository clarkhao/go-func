package linked

import (
	"fmt"
)

// Definition for singly-linked list.
type LinkedList[T interface{}] struct {
	Value T
	Next  *LinkedList[T]
}

// NewNode is a function that create a new node of LinkedList
// return the new node's address
func NewNode[T interface{}](value T) *LinkedList[T] {
	return &(LinkedList[T]{
		Value: value,
		Next:  nil,
	})
}

// Add is a method that add a new node with value to the end of LinkedList
// return the address of new added node
func (node *LinkedList[T]) Push(value T) *LinkedList[T] {
	new := NewNode[T](value)
	if node == nil {
		return new
	}
	(*node).Next = new
	return (*node).Next
}

// IterateWithCh is a method that iterate all values in LinkedList
func (node *LinkedList[T]) IterateWithCh() (<-chan T, func()) {
	ch := make(chan T)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		defer close(ch)
		current := node
		for current != nil {
			select {
			case ch <- current.Value:
				current = current.Next
			case <-done:
				fmt.Println("done here")
				return
			}
		}
	}()
	return ch, cancel
}

// ToList is a method that turn a LinkedList to a list
// return the list
func (node *LinkedList[T]) ToList() (list []T) {
	ch, cancel := node.IterateWithCh()
	for value := range ch {
		list = append(list, value)
	}
	cancel()
	return list
}

// Len is a method that count the length of LinkedList
// return int
func (node *LinkedList[T]) Len() int {
	if node == nil {
		return 0
	}
	return node.Next.Len() + 1
}
