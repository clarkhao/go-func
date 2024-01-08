package stack_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/clarkhao/go-func/stack"
)

func TestSafePush(t *testing.T) {
	cases := []struct {
		inputStack []int
		outputSize int
	}{
		{
			inputStack: []int{1, 2, 3, 4, 5},
			outputSize: 5,
		},
	}
	for _, c := range cases {
		var wg sync.WaitGroup
		s := stack.SafeStack[int]{
			&stack.Stack[int]{},
			sync.RWMutex{},
		}
		for _, v := range c.inputStack {
			wg.Add(1)
			go s.Push(v, &wg)
		}
		wg.Wait()
		length := s.Len()
		list := []int{}
		for v := range s.Iterate() {
			list = append(list, v)
		}
		fmt.Println(list)
		if len(list) != c.outputSize && length != c.outputSize {
			t.Errorf("failed to iterate stack")
		}
	}
}

func TestSafePop(t *testing.T) {
	cases := []struct {
		inputStack []int
		outputSize int
	}{
		{
			inputStack: []int{1, 2, 3, 4, 5},
			outputSize: 3,
		},
		{
			inputStack: []int{1, 2},
			outputSize: 0,
		},
	}
	for _, c := range cases {
		var wg sync.WaitGroup
		s := stack.SafeStack[int]{
			&stack.Stack[int]{},
			sync.RWMutex{},
		}
		for _, v := range c.inputStack {
			wg.Add(1)
			go s.Push(v, &wg)
		}
		wg.Wait()
		for i := 0; i < 2; i++ {
			wg.Add(1)
			s.Pop(&wg)
		}
		wg.Wait()
		length := s.Len()
		list := []int{}
		for v := range s.Iterate() {
			list = append(list, v)
		}
		fmt.Println(list)
		if len(list) != c.outputSize && length != c.outputSize {
			t.Errorf("failed to Pop stack")
		}
	}
}
