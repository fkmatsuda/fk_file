package fkfile

import (
	"archive/zip"
	"bufio"
	"compress/gzip"
	"io"
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

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}

	w := gzip.NewWriter(dstFile)
	defer w.Close()

	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}

	return w.Flush()

}

// ZipFile defines a zip file
type ZipFile struct {
	file   *os.File
	writer *zip.Writer
}

// NewZipFile creates a new zip file
func NewZipFile(file *os.File) (zipFile *ZipFile, err error) {
	zipFile = &ZipFile{
		file:   file,
		writer: zip.NewWriter(file),
	}
	return
}

// Close closes the zip file
func (zipFile *ZipFile) Close() error {
	err := zipFile.writer.Close()
	if err != nil {
		return err
	}
	return zipFile.file.Close()
}

// AddEntry adds a file to the zip file
func (zipFile *ZipFile) AddEntry(name string, reader io.Reader) error {
	w, err := zipFile.writer.Create(name)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, reader)
	if err != nil {
		return err
	}
	return nil
}
