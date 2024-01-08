package linked_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/clarkhao/go-func/linked"
)

func TestNewNode(t *testing.T) {
	cases := []struct {
		input    string
		expected linked.LinkedList[string]
	}{
		{input: "Hello", expected: linked.LinkedList[string]{Value: "Hello", Next: nil}},
	}
	for _, c := range cases {
		output := linked.NewNode[string](c.input)
		if output.Value != c.expected.Value {
			t.Errorf("NewNode(%v) = %v, expected: %v", c.input, output, c.expected)
		}
	}
}

func TestPush(t *testing.T) {
	cases := []struct {
		input    string
		expected *linked.LinkedList[string]
	}{
		{input: "Hello", expected: &(linked.LinkedList[string]{Value: "Hello", Next: nil})},
		{input: "World", expected: linked.NewNode[string]("Hello").Push("World")},
		{input: "I am Clark", expected: linked.NewNode[string]("Hello").Push("World").Push("I am Clark")},
	}
	var node *linked.LinkedList[string]
	var head = node
	for i, c := range cases {
		fmt.Println(node)
		node = node.Push(c.input)
		if i == 0 {
			head = node
		}
		fmt.Println(head)
		if node.Value != c.expected.Value {
			t.Errorf("node.Push(%v) = %v, expected %v", c.input, node, c.expected)
		}
	}
}

func TestIterateWithCh(t *testing.T) {
	cases := []struct {
		expected []string
	}{
		{expected: []string{"Hello", "World", "I", "Am", "Clark"}},
	}
	for _, c := range cases {
		var node *linked.LinkedList[string] = nil
		head := node
		for i := 0; i < len(c.expected); i++ {
			node = node.Push(c.expected[i])
			if i == 0 {
				head = node
			}
		}
		ch, cancel := head.IterateWithCh()
		outputSlice := []string{}
		for value := range ch {
			outputSlice = append(outputSlice, value)
		}

		fmt.Println("Loop Ended")
		if !slices.Equal[[]string, string](c.expected, outputSlice) {
			t.Errorf("Iterate() produce %v, expected %v", outputSlice, c.expected)
		}
		cancel()
	}
}

func TestToList(t *testing.T) {
	cases := []struct {
		expected []string
	}{
		{expected: []string{"Hello", "World", "I", "Am", "Clark"}},
		{expected: []string{}},
	}
	for _, c := range cases {
		var node *linked.LinkedList[string] = nil
		head := node
		for i := 0; i < len(c.expected); i++ {
			node = node.Push(c.expected[i])
			if i == 0 {
				head = node
			}
		}
		fmt.Println(head)
		output := head.ToList()
		if !slices.Equal[[]string, string](c.expected, output) {
			t.Errorf("Iterate() produce %v, expected %v", output, c.expected)
		}
	}
}

func TestLen(t *testing.T) {
	cases := []struct {
		expected int
	}{
		{expected: 10},
		{expected: 0},
		{expected: 3},
	}

	for _, c := range cases {
		var node *linked.LinkedList[int] = nil
		head := node
		for i := 0; i < c.expected; i++ {
			node = node.Push(i)
			if i == 0 {
				head = node
			}
		}
		output := head.Len()
		if output != c.expected {
			t.Errorf("length of linkedlist is %v, expected %v", output, c.expected)
		}
	}
}
