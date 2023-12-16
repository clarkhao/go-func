package cp

import (
	"fmt"
	"os"
	"path/filepath"
)

// CopyOne is a function which copy file from src to dest
// return error if failed
func CopyOne(src string, dest string) error {
	source, err := os.Open(src)
	if err != nil {
		err = fmt.Errorf("in Copyone when openning %w", err)
		return err
	}
	defer source.Close()
	destination, err := os.Create(dest)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File not found:", dest)
		}
		err = fmt.Errorf("in Copyone when creating destination %w", err)
		return err
	}
	defer destination.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := source.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				err = fmt.Errorf("in Copyone when creating buf %w", err)
				return err
			}
			break
		}
		_, err = destination.Write(buffer[:n])
		if err != nil {
			err = fmt.Errorf("in Copyone when writing %w", err)
			return err
		}
	}
	return nil
}

// CopyDir is a function which copy all files and dirs within src and paste to dest
// return error if failed
func CopyDir(src string, dest string) error {
	err := os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		err = fmt.Errorf("in CopyDir when mkdir dest with error %w", err)
		return err
	}
	files, err := os.ReadDir(src)
	if err != nil {
		err = fmt.Errorf("in CopyDir when read all files and dirs inside src with error %w", err)
		return err
	}
	for _, file := range files {
		srcPath := filepath.Join(src, file.Name())
		destPath := filepath.Join(dest, file.Name())
		if file.IsDir() {
			err = CopyDir(srcPath, destPath)
			if err != nil {
				err = fmt.Errorf("in CopyDir when recursive copy with error %w", err)
				return err
			}
		} else {
			err = CopyOne(srcPath, destPath)
			if err != nil {
				err = fmt.Errorf("in CopyDir when copy one file failed with error %w", err)
				return err
			}
		}
	}
	return nil
}
