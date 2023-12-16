package cp_test

import (
	"testing"

	"github.com/clarkhao/go-func/cp"
)

func TestCopyOne(t *testing.T) {
	cases := []struct {
		inputSrc  string
		inputDest string
		outputErr bool
	}{
		{"C:\\tmp\\abc.txt", "C:\\tmp\\a.txt", true},
	}
	for _, c := range cases {
		err := cp.CopyOne(c.inputSrc, c.inputDest)
		if err != nil {
			t.Errorf("CopyOne(%v, %v) = %v, expected %v", c.inputSrc, c.inputDest, err, c.outputErr)
		}
	}
}
func TestCopyDir(t *testing.T) {
	cases := []struct {
		inputSrc  string
		inputDest string
		outputErr bool
	}{
		{"C:\\tmp", "C:\\atmp", true},
	}
	for _, c := range cases {
		err := cp.CopyDir(c.inputSrc, c.inputDest)
		if err != nil {
			t.Errorf("CopyOne(%v, %v) = %v, expected %v", c.inputSrc, c.inputDest, err, c.outputErr)
		}
	}
}
