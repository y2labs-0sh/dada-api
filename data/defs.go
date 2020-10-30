package data

import (
	"os"
	"path/filepath"
)

func GetFileModTime(path string) (int64, error) {

	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return 0, err
	}
	defer (*f).Close()

	fi, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return fi.ModTime().Unix(), nil
}
