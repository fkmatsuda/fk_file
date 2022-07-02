package fkfile

import (
	"io/ioutil"
	"os"
	"testing"
)

func setup() string {
	dir, err := ioutil.TempDir("", "*")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(dir+"/test.txt", []byte("Hello, test!"), 0644)
	if err != nil {
		panic(err)
	}
	return dir
}

func TestFile(t *testing.T) {
	tempDir := setup()
	defer os.RemoveAll(tempDir)
	t.Run("ReadTextFile", func(t *testing.T) {
		text, err := ReadTextFile(tempDir + "/test.txt")
		if err != nil {
			t.Error(err)
		}
		if text != "Hello, test!" {
			t.Error("Expected 'Hello, test!'")
		}
	})
	t.Run("EnsureDir", func(t *testing.T) {
		dirPath := tempDir + "/test"
		err := EnsureDir(dirPath)
		if err != nil {
			t.Error(err)
		}
		if _, err := os.Stat(dirPath); err != nil {
			if os.IsNotExist(err) {
				t.Error("Expected directory")
			} else {
				t.Error(err)
			}
		}
	})
}
