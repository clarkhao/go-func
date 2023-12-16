package clean_test

import (
	"testing"

	"github.com/clarkhao/go-func/clean"
)

func TestDirExist(t *testing.T) {
	cases := []struct {
		input  string
		output bool
	}{
		{"", false},
		{"C:\\", true},
		{"C:\\Users", true},
	}
	for _, c := range cases {
		isExist, err := clean.DirExist(c.input)
		if err != nil {
			t.Errorf("DirExist(%v) = %v, expected: %v", c.input, isExist, c.output)
		}
	}
}

func TestRemoveWithin(t *testing.T) {
	cases := []struct {
		input  string
		output bool
	}{
		{"C:\\tmp", false},
	}
	for _, c := range cases {
		err := clean.RemoveWithin(c.input)
		if err != nil {
			t.Errorf("RemoveRecursive(%v) = %v, expected: %v", c.input, err.Error(), c.output)
		}
	}
}
