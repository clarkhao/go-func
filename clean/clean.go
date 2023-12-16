package clean

import (
	"os"
	"path"
	"strings"
)

// DirExist is a function that determine if dir exists or not
// return bool & error
func DirExist(dir string) (bool, error) {
	info, err := os.Stat(dir)
	if err == nil {
		return info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// RemoveWithin is a function which delete all files and dirs within root dir specified
// return error if failed but nil if succeeded
func RemoveWithin(root string) error {
	files, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	for _, file := range files {

		if !(strings.Contains(file.Name(), ".git")) {
			if file.IsDir() {
				os.RemoveAll(path.Join(root, file.Name()))
			} else {
				os.Remove(path.Join(root, file.Name()))
			}
		}
	}
	return nil
}
