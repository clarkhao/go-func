package queue_test

import (
	"fmt"
	"slices"
	"sync"
	"testing"

	"github.com/clarkhao/go-func/queue"
)

func TestPush(t *testing.T) {
	cases := []struct {
		inputQueue []int
		outputSize int
	}{
		{
			inputQueue: []int{1, 2, 3, 4, 5},
			outputSize: 5,
		},
	}
	for _, c := range cases {
		var wg sync.WaitGroup
		q := queue.Queue[int]{
			&queue.UnsafeQueue[int]{},
			sync.RWMutex{},
		}
		for _, v := range c.inputQueue {
			wg.Add(1)
			go q.Push(v, &wg)
		}
		wg.Wait()
		ch := q.Queue.IterateFromHead()
		var headList []int
		for v := range ch {
			headList = append(headList, v)
		}
		fmt.Println(headList)
		ch = q.Queue.IterateFromTail()
		var tailList []int
		for v := range ch {
			tailList = append(tailList, v)
		}
		fmt.Println(tailList)
		slices.Reverse[[]int](tailList)
		if len(headList) != c.outputSize && !slices.Equal[[]int](tailList, headList) {
			t.Errorf("failed to safe push Queue")
		}
	}
}

func TestPop(t *testing.T) {
	cases := []struct {
		inputQueue []int
	}{
		{inputQueue: []int{1, 2, 3, 4, 5}},
	}
	for _, c := range cases {
		var wg sync.WaitGroup
		q := queue.Queue[int]{
			&queue.UnsafeQueue[int]{},
			sync.RWMutex{},
		}
		for _, v := range c.inputQueue {
			wg.Add(1)
			go q.Push(v, &wg)
		}
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go q.Pop(&wg)
		}
		wg.Wait()
		ch := q.Queue.IterateFromHead()
		var headList []int
		for v := range ch {
			headList = append(headList, v)
		}
		fmt.Println(headList)
		ch = q.Queue.IterateFromTail()
		var tailList []int
		for v := range ch {
			tailList = append(tailList, v)
		}
		fmt.Println(tailList)
		slices.Reverse[[]int](tailList)
		if !slices.Equal[[]int](tailList, headList) {
			t.Errorf("failed to safe pop Queue")
		}
	}
}
