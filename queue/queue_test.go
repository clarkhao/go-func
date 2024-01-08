package queue_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/clarkhao/go-func/queue"
)

func TestUnsafePush(t *testing.T) {
	cases := []struct {
		inputQueue     []int
		outputHeadList []int
		outputTailList []int
	}{
		{
			inputQueue:     []int{1, 2, 3, 4},
			outputTailList: []int{4, 3, 2, 1},
		},
	}
	for _, c := range cases {
		uq := queue.UnsafeQueue[int]{}
		for _, v := range c.inputQueue {
			uq.Push(v)
		}
		fmt.Println(uq.Len())
		var headList []int
		for v := range uq.IterateFromHead() {
			headList = append(headList, v)
		}
		fmt.Println(headList)
		var tailList []int
		for v := range uq.IterateFromTail() {
			tailList = append(tailList, v)
		}
		fmt.Println(tailList)
		if !(slices.Equal[[]int](c.inputQueue, headList) && slices.Equal[[]int](c.outputTailList, tailList)) {
			t.Errorf("failed to push or iterate Queue")
		}
	}
}

func TestUnsafePop(t *testing.T) {
	cases := []struct {
		inputQueue     []int
		outputHeadList []int
		outputTailList []int
		popValid       bool
	}{
		{
			inputQueue:     []int{1, 2, 3, 4},
			outputHeadList: []int{3, 4},
			outputTailList: []int{4, 3},
			popValid:       true,
		},
		{
			inputQueue:     []int{1},
			outputHeadList: []int{},
			outputTailList: []int{},
			popValid:       false,
		},
	}
	for _, c := range cases {
		uq := queue.UnsafeQueue[int]{}
		for _, v := range c.inputQueue {
			uq.Push(v)
		}
		fmt.Printf("after push the length is %v\n", uq.Len())
		uq.Pop()
		_, ok := uq.Pop()
		fmt.Printf("after pop the length is %v\n", uq.Len())
		var headList []int
		for v := range uq.IterateFromHead() {
			headList = append(headList, v)
		}
		fmt.Println(headList)
		var tailList []int
		for v := range uq.IterateFromTail() {
			tailList = append(tailList, v)
		}
		fmt.Println(tailList)
		if !(slices.Equal[[]int](c.outputHeadList, headList) && slices.Equal[[]int](c.outputTailList, tailList) && ok == c.popValid) {
			t.Errorf("failed to push or iterate Queue")
		}
	}
}
