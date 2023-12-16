package search_test

import (
	"testing"

	"github.com/clarkhao/go-func/search"
	"github.com/clarkhao/go-func/slices"
)

func TestSearchFile(t *testing.T) {
	cases := []struct {
		inputSrc  string
		intpuDest string
		expected  []string
	}{
		{"C:\\tmp", "a.txt", []string{"C:\\tmp\\dddsr", "C:\\tmp"}},
	}
	for _, c := range cases {
		output, _ := search.SearchFile(c.inputSrc, c.intpuDest)
		outputSlice := slices.NewSlice(output)
		if !outputSlice.ItemEqual(c.expected) {
			t.Errorf("SearchFile(%v, %v) = %v, expected: %v", c.inputSrc, c.intpuDest, output, c.expected)
		}
	}
}
