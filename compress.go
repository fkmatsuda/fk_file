package fkfile

import (
	"bufio"
	"compress/gzip"
	"io/ioutil"
	"os"
)

// Compress a file with gz algorithm
func CompressFile(src, dst string, removeSrc bool) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	r := bufio.NewReader(srcFile)

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}

	w := gzip.NewWriter(dstFile)
	defer w.Close()

	_, err = w.Write(data)
	if err != nil {
		return err
	}

	return w.Flush()

}
