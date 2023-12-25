package json_test

import (
	"fmt"
	"testing"

	"github.com/clarkhao/go-func/json"
)

func TestReadJson(t *testing.T) {
	cases := []struct {
		inputFile string
		expected  bool
	}{
		{
			inputFile: "C:\\Users\\clark\\Documents\\codes\\ts\\vite-function\\package.json",
			expected:  false,
		},
	}
	for _, c := range cases {
		jsonItems, err := json.ReadJson(c.inputFile)
		if err != nil {
			t.Errorf("json.ReadJson(%v) = %v, expected without error %v", c.inputFile, err, c.expected)
		}
		fmt.Println(*jsonItems)
	}
}

func TestAddItems(t *testing.T) {
	cases := []struct {
		inputTarget  string
		inputExtense string
		expected     int
	}{
		{
			inputTarget:  "C:\\Users\\clark\\Documents\\codes\\ts\\vite-function\\package.json",
			inputExtense: "C:\\Users\\clark\\Documents\\codes\\ts\\vite-function\\some.json",
			expected:     1,
		},
	}
	for _, c := range cases {
		num, err := json.AddItems(c.inputTarget, c.inputExtense)
		if err != nil || num != c.expected {
			t.Errorf("AddItems(%v, %v) = %v, expected %v", c.inputTarget, c.inputExtense, num, c.expected)
		}
	}
}
