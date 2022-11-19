package xfile

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetCurrentDirectory ...
func GetCurrentDirectory() (dir string, err error) {
	if dir, err = os.Getwd(); err == nil {
		return
	}

	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("err", err)
	}

	dir = strings.Replace(dir, "\\", "/", -1)
	return
}

// MakeDirectory ...
func MakeDirectory(path string) (err error) {
	if err = Exists(path); err != nil {
		return os.MkdirAll(path, 0755)
	}

	return
}

// Exists ...
func Exists(path string) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return
	}

	return
}
