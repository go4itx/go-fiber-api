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
	if !Exists(path) {
		return os.MkdirAll(path, 0755)
	}

	return
}

// Exists ...
func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// log.Println(path, err)
		return false
	}

	return true
}

// ListFiles returns all file names in `dir`
func ListFiles(dir string, ext string) (res []string) {
	res = make([]string, 0)
	fs, err := os.ReadDir(dir)
	if err != nil {
		return res
	}

	for _, fp := range fs {
		if fp.IsDir() {
			continue
		}

		if ext != "" && filepath.Ext(fp.Name()) != ext {
			continue
		}

		res = append(res, fp.Name())
	}

	return res
}
