package fkfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func setupCompress() string {
	dir, err := ioutil.TempDir("", "*")
	if err != nil {
		panic(err)
	}
	testString := "Hello, test!"
	for i := 0; i < 100; i++ {
		testString = fmt.Sprintf("%s\n", testString)
	}
	err = os.WriteFile(dir+"/test.txt", []byte(testString), 0644)
	if err != nil {
		panic(err)
	}
	return dir
}

func TestCompress(t *testing.T) {
	tempDir := setupCompress()
	defer os.RemoveAll(tempDir)
	t.Run("CompressFile", func(t *testing.T) {
		filePath := tempDir + "/test.txt"
		compressPath := tempDir + "/test.txt.gz"
		err := CompressFile(filePath, tempDir+"/test.txt.gz", true)
		if err != nil {
			t.Error(err)
		}
		if _, err := os.Stat(compressPath); err != nil {
			if os.IsNotExist(err) {
				t.Error("Expected compressed file")
			} else {
				t.Error(err)
			}
		}
		fi, err := os.Stat(filePath)
		if err != nil {
			t.Error(err)
		}
		cfi, err := os.Stat(compressPath)
		if err != nil {
			t.Error(err)
		}
		if cfi.Size() > fi.Size() {
			t.Error("Expected compressed file to be smaller")
		}
		if cfi.Size() == 0 {
			t.Error("Expected compressed file to be non-empty")
		}
	})
	t.Run("Open source file error", func(t *testing.T) {
		filePath := tempDir + "/notExists.txt"
		err := CompressFile(filePath, tempDir+"/test.txt.gz", true)
		if err == nil {
			t.Error("Expected error")
		}
		if !os.IsNotExist(err) {
			t.Error("Expected the error to be NotExists")
		}
	})
}
