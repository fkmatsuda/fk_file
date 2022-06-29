package fkfile

import (
	"io/ioutil"
)

// WriteTextFile writes text to a file
func WriteTextFile(filename string, text string) (err error) {
	err = ioutil.WriteFile(filename, []byte(text), 0644)
	return
}

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
	} else {
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
			} else {
				return err
			}
		}
	}
	return nil
}
