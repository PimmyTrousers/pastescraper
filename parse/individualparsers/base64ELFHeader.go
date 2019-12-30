package individualparsers

import (
	"bytes"
	"io/ioutil"
)

type Base64ELFHeader struct {}

func (b Base64ELFHeader) Match(content []byte) (bool, error) {
	buf, err := ioutil.ReadAll(bytes.NewBuffer(content))
	if err != nil {
		return false, err
	}

	// ELF header base64 encoded
	if string(buf[:5]) == "f0VMR" {
		return true, nil
	}

	return false, nil
}
