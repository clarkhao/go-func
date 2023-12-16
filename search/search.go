package search

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

// SearchFile is a function which search for all destFile inside srcPath dir
// return all dirs including the destFile or error
func SearchFile(srcPath string, destFile string) (paths []string, err error) {
	filepath.Walk(srcPath, func(currentPath string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		d := filepath.Dir(currentPath)
		f := filepath.Base(currentPath)
		if f == destFile {
			paths = append(paths, d)
		}
		return nil
	})
	if err != nil {
		err = fmt.Errorf("failed to walk directory: %w", err)
		return
	}
	return
}
