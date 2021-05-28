package fkfile

import (
	"io/ioutil"
)

// ReadTextFile returns text of a file
func ReadTextFile(filename string) (text string, err error) {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	text = string(bytes)

	return

}
