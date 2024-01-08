package stack_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/clarkhao/go-func/stack"
)

func TestIterate(t *testing.T) {
	cases := []struct {
		inputStack  []int
		outputStack []int
	}{
		{
			inputStack:  []int{1, 2, 3},
			outputStack: []int{3, 2, 1},
		},
		{
			inputStack:  []int{},
			outputStack: []int{},
		},
	}
	for _, c := range cases {
		var s stack.Stack[int]
		for _, v := range c.inputStack {
			s.Push(v)
		}
		list := []int{}
		for v := range s.Iterate() {
			list = append(list, v)
		}
		if !(slices.Equal[[]int](list, c.outputStack)) {
			t.Error("failed to iterate the stack")
		}
	}
}

func TestPop(t *testing.T) {
	cases := []struct {
		inputStack    []int
		afterPopStack []int
		popValid      bool
	}{
		{
			inputStack:    []int{1, 2, 3},
			afterPopStack: []int{1},
			popValid:      true,
		},
		{
			inputStack:    []int{},
			afterPopStack: []int{},
			popValid:      false,
		},
	}
	for _, c := range cases {
		var s stack.Stack[int]
		for _, v := range c.inputStack {
			s.Push(v)
		}
		s.Pop()
		_, ok := s.Pop()
		fmt.Printf("%v\n", ok)
		list := []int{}
		for v := range s.Iterate() {
			list = append(list, v)
		}
		if !(slices.Equal[[]int](list, c.afterPopStack)) || ok != c.popValid {
			t.Error("failed to Pop the stack")
		}
	}
}
