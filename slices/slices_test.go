package slices_test

import (
	"testing"

	"github.com/clarkhao/go-func/slices"
)

func TestToMap(t *testing.T) {
	cases := []struct {
		inputList slices.Slices[[]string, string]
		expected  map[string]int
	}{
		{
			slices.Slices[[]string, string]{"a", "b", "c"},
			map[string]int{"a": 1, "b": 1, "c": 1},
		},
	}
	for _, c := range cases {
		output := c.inputList.ToMap()
		if len(output) != len(c.expected) {
			t.Errorf("%v.ToMap() = %v, expected: %v", c.inputList, output, c.expected)
		}
		for key, value := range output {
			if c.expected[key] != value {
				t.Errorf("%v.ToMap() = %v, expected: %v", c.inputList, output, c.expected)
			}
		}
	}
}

func TestItemEqual(t *testing.T) {
	cases := []struct {
		inputListObj slices.Slices[[]string, string]
		inputSlice   slices.Slices[[]string, string]
		expected     bool
	}{
		{
			slices.Slices[[]string, string]{"a", "b", "a"},
			slices.Slices[[]string, string]{"b", "a", "a"},
			true,
		},
	}
	for _, c := range cases {
		output := c.inputListObj.ItemEqual(c.inputSlice)
		if output != c.expected {
			t.Errorf("%v.ItemEqual(%v) = %v, expected %v", c.inputListObj, c.inputSlice, output, c.expected)
		}
	}
}
