package fkfile

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadTextFile returns text of a file
func ReadTextFile(filename string) (text string, err error) {
	bytes, err := ioutil.ReadFile(filename)
	if err == nil {
		text = string(bytes)
	}
	return
}

// EnsureDir ensures that the directory exists
func EnsureDir(path string) error {
	parent := filepath.Dir(path)
	if parent == path {
		return nil
	}
	err := EnsureDir(parent)
	if err != nil {
		return err
	}
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		}
		return err
	}
	return nil
}
